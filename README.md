# Steps to follow in order to run the code

Note: Open the src folder as the root of the project

################################################# Clone raft ##########################


In order to run the code first we need to clone raft, to do that run the below command:

1. Go to github.com folder
cd src/github.com

2. If go.etcd.io folder is not there run the below command
mkdir go.etcd.io

3. Go inside go.etcd.io
cd go.etcd.io

4. Inside this folder run:
git clone https://github.com/etcd-io/etcd.git

5. Go to the src by running the below commands
cd ..
cd ..

6. Go to src folder and run this command
go mod tidy

7. Once the above step and clone is done go inside raftexample folder

cd github.com/go.etcd.io/etcd/contrib/raftexample

8. And do build of raftexample by running the below command:
go build -o raftexample

9. And run the below command
goreman start

################################################# Start Services ######################


Open four other terminals under src and run the following commands in the respective terminals: 

1st terminal
cd web/services/tokenActions/tokenActions_service
go run tokenActionService.go

2nd terminal
cd web/services/user/user_service
go run userService.go

3rd terminal
cd web/services/post/post_service
go run postService.go

4th terminal
cd cmd/web
go run web.go

Once above steps are done go to browser and open this URL
http://localhost:8000/login/

### Note: If any error occurs stop the goreman service, run `goreman start` and restart the services to make sure services and goreman started properly ###

################################################# Running Test Cases ###################

Order of running the services:
1. goreman start, then wait for goreman to display that a node has been selected as a leader
2. tokenActionsService
3. userService
4. postService
5. web.go

### Note: If any error occurs or test cases don't run, stop the goreman service, run `goreman start`, wait till a leader has been elected and then restart the services to make sure services and goreman started properly ###

1. Stop all services and goreman (ctrl + c) in all 5 terminals
2. Open another terminal and go to: cd web/testCases/
3. Start goreman and then start all 4 services
4. Run the test cases by running the following (as one command): go run mainTest.go registerLoginTest.go followUnfollowTest.go postTest.go

### Note: To re-run test cases, please stop all services and goreman and then restart in the order mentioned above (goreman first, then wait for leader to be elected in goreman and then start the 4 services), this is to flush all objects out of raft and re-connect to the services ###
