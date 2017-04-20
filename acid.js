// Acid CI/CD

// We need to simulate a Go environment
gopath = "/go";
localPath = gopath + "/src/github.com/" + pushRecord.repository.full_name;

// Define a single build step:
j = new Job("unit-tests");
j.image = "golang:1.7.5";
j.env = {
  "DEST_PATH": localPath,
  "GOPATH": gopath
};

j.tasks = [
	"echo 'running tests'",
  "make bootstrap",
  "make test"
];

j.run(pushRecord).waitUntilDone()
