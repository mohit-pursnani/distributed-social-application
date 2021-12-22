package main

import (
	"context"
	"log"
	"sync"

	"github.com/os3224/final-project-0b5a2e16-santhoshbhk-mohit-pursnani/web/services/post/post_pb"
)

func BatchPostTest() {
	log.Println("xxxxxxxxxxRunning Concurrent Post Test casexxxxxxxxxx")
	passedTestCases := true

	wg := sync.WaitGroup{}
	for i := 0; i < 1000; i += 1 {
		wg.Add(1)

		go func(i int) {
			defer wg.Done()
			post := post_pb.Post{
				UserID: int32(i), PostText: "testPost",
			}

			_, err := clientObj.PostClient.AddPost(context.Background(), &post)
			if err != nil {
				log.Fatal("Error: ", err)
				passedTestCases = false
			}
		}(i)
	}

	wg.Wait()
	post := post_pb.Post{
		UserID: int32(500), PostText: "testPost",
	}

	postLists, _ := clientObj.PostClient.AddPost(context.Background(), &post)

	if len(postLists.PostList) != 1001 {
		passedTestCases = false
	}

	if passedTestCases {
		log.Println("Posts Test cases Passed!")
	} else {
		log.Println("Posts Test cases failed...")
	}

}
