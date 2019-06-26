echo 'Running tets...'
./run_tests.sh
echo 'Building Zoroaster...'
env GOOS=linux GOARCH=amd64 go build -o zoroaster
echo 'Uploading binary to ec2...'
scp zoroaster ec2-user@ec2-3-121-239-182.eu-central-1.compute.amazonaws.com:.
echo 'Done'
