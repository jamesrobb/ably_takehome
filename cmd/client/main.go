package main

import (
	"context"
	"crypto/md5"
	"encoding/hex"
	"flag"
	"fmt"
	"io"
	"math/rand"
	"os"
	"time"

	"github.com/google/uuid"

	"github.com/jamesrobb/ably-takehome/protocol/generated/protocol"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

const MAX_NUMBERS uint32 = 65535

func main() {
	port := flag.Int("port", 50051, "port the of the server to be connected to")
	numMessagesFlag := flag.Uint("numMessages", 0, "number of messages to receive, specifying 0 will result in a random value beteen 1 and 65535")
	testUUID := flag.String("testUUID", "", "UUID used (used in test mode only)")
	testChecksum := flag.String("testChecksum", "", "expected checksum of successful result (used in test mode only)")
	seed := flag.Uint("testSeed", 1, "seed used for the server's PRNG (used in test mode only)")
	testMode := flag.Bool("testMode", false, "run a sanity check on an interrupted stream")
	flag.Parse()

	serverAddress := fmt.Sprintf("localhost:%d", *port)

	numNumbers := uint32(*numMessagesFlag)
	if numNumbers == 0 {
		rand.Seed(time.Now().Unix())
		numNumbers = uint32(rand.Intn(int(MAX_NUMBERS)-1)) + 1
	}
	if numNumbers > MAX_NUMBERS {
		numNumbers = MAX_NUMBERS
	}

	if *testMode {
		uuid, err := uuid.Parse(*testUUID)
		if err != nil {
			fmt.Println("FAILURE: unable to parse provided UUID")
			os.Exit(1)
		}
		err = testOperation(serverAddress, numNumbers, uuid, uint32(*seed), *testChecksum)
		if err != nil {
			fmt.Printf("FAILURE: %s\n", err)
			os.Exit(1)
		}

		return
	} else {
		err := standardOperation(serverAddress, numNumbers)
		if err != nil {
			fmt.Printf("FAILURE: %s\n", err)
			os.Exit(1)
		}

		return
	}
}

func getClient(serverAddress string) (*grpc.ClientConn, protocol.NumbersClient, error) {
	var opts []grpc.DialOption

	opts = append(opts, grpc.WithTransportCredentials(insecure.NewCredentials()))
	opts = append(opts, grpc.WithBlock())

	conn, err := grpc.Dial(serverAddress, opts...)
	if err != nil {
		return nil, nil, fmt.Errorf("unable to dial into server\n")
	}

	return conn, protocol.NewNumbersClient(conn), nil
}

func testOperation(
	serverAddress string,
	numMessages uint32,
	uuid uuid.UUID,
	seed uint32,
	testChecksum string,
) error {
	if numMessages%2 != 0 {
		return fmt.Errorf("for testMode specify an even number of messages to be received")
	}

	conn1, client1, err := getClient(serverAddress)
	if err != nil {
		return err
	}

	numbers1, _, err := getNumbers(client1, uuid, numMessages, seed, numMessages/2)
	if err != nil {
		return fmt.Errorf("error getting first batch of numbers: %s", err)
	}
	conn1.Close()

	time.Sleep(2 * time.Second)

	_, client2, err := getClient(serverAddress)
	if err != nil {
		return err
	}
	numbers2, serverChecksum, err := getNumbers(client2, uuid, numMessages, 0, 0)
	if err != nil {
		return fmt.Errorf("error getting second batch of numbers: %s", err)
	}

	for _, num := range numbers2 {
		numbers1 = append(numbers1, num)
	}

	calculatedChecksum := calculateChecksum(numbers1)
	if calculatedChecksum != serverChecksum {
		return fmt.Errorf("calculatedChecksum=%s does not match serverChecksum=%s\n", calculatedChecksum, serverChecksum)
	}
	if testChecksum != calculatedChecksum {
		return fmt.Errorf("testChecksum=%s does not match calculatedChecksum=%s\n", testChecksum, calculatedChecksum)
	}

	fmt.Printf("SUCCESS: checksum=%s\n", serverChecksum)

	return nil
}

func standardOperation(serverAddress string, numMessages uint32) error {
	_, client, err := getClient(serverAddress)
	if err != nil {
		return err
	}

	u := uuid.New()
	numbers, serverChecksum, err := getNumbers(client, u, numMessages, 0, 0)
	if err != nil {
		return fmt.Errorf("error getting numbers: %s\n", err)
	}

	calculatedChecksum := calculateChecksum(numbers)
	if calculatedChecksum != serverChecksum {
		return fmt.Errorf("calculatedChecksum=%s does not match serverChecksum=%s\n", calculatedChecksum, serverChecksum)
	}

	fmt.Printf("success checksum=%s\n", serverChecksum)

	return nil
}

func calculateChecksum(numbers []uint32) string {
	h := md5.New()

	for _, num := range numbers {
		io.WriteString(h, fmt.Sprintf("%d", num))
	}

	return hex.EncodeToString(h.Sum(nil)[:])
}

func getNumbers(
	client protocol.NumbersClient,
	clientUUID uuid.UUID,
	numNumbers uint32,
	seed uint32,
	breakAfter uint32,
) ([]uint32, string, error) {
	m := &protocol.NumbersRequest{
		ClientId:   clientUUID[:],
		NumNumbers: numNumbers,
		Seed:       seed,
	}

	stream, err := client.GetNumbers(context.Background(), m)
	if err != nil {
		return nil, "", err
	}

	serverChecksum := ""
	numbers := make([]uint32, 0)

	for {
		number, err := stream.Recv()
		if err == io.EOF {
			break
		} else if err != nil {
			return nil, "", fmt.Errorf("error reading from stream: %s", err)
		}

		fmt.Println(number.Number)
		numbers = append(numbers, number.Number)

		// if we have a non-empty checksum then the number stream is finished
		if number.Checksum != "" {
			serverChecksum = number.Checksum
			stream.CloseSend()

			break
		}

		// breakAfter is to be able to simulate a connection being broken
		if breakAfter > 0 {
			if len(numbers) == int(breakAfter) {
				stream.CloseSend()

				return numbers, "", nil
			}
		}
	}

	return numbers, serverChecksum, nil
}
