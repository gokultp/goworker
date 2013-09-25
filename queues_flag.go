package goworker

import (
	"errors"
	"fmt"
	"strconv"
	"strings"
)

var (
	errorEmptyQueues      = errors.New("You must specify at least one queue.")
	errorNonNumericWeight = errors.New("The weight must be a numeric value.")
)

type queuesFlag []string

func (q *queuesFlag) Set(value string) error {
	// Parse the individual queues and their weights if they are present.
	for _, queueAndWeight := range strings.Split(value, ",") {
		if queueAndWeight == "" {
			continue
		}

		queue, weight, err := parseQueueAndWeight(queueAndWeight)
		if err != nil {
			return err
		}

		for i := 0; i < weight; i++ {
			*q = append(*q, queue)
		}
	}
	if len(*q) == 0 {
		return errorEmptyQueues
	}
	return nil
}

func (q *queuesFlag) String() string {
	return fmt.Sprint(*q)
}

func parseQueueAndWeight(queueAndWeight string) (queue string, weight int, err error) {
	parts := strings.SplitN(queueAndWeight, "=", 2)
	queue = parts[0]

	if queue == "" {
		return
	}

	// If '=' is not present then we only have the queue name and the default weight is 1.
	if len(parts) == 1 {
		weight = 1
		err = nil
		return
	}

	// Check to see if we have a weight for this queue.
	if len(parts) == 2 {
		weight, err = strconv.Atoi(parts[1])
		if err != nil {
			return "", 0, errorNonNumericWeight
		}
		return queue, weight, nil
	}
	return
}
