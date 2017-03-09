package worker

import (
	"tes/config"
	pbe "tes/ga4gh"
	"tes/scheduler"
	"tes/server"
	server_mocks "tes/server/mocks"
	pbr "tes/server/proto"
)

// mockScheduler is a mock scheduler that assigns every job to a single worker.
type mockScheduler struct {
	worker *pbr.Worker
}

func (m *mockScheduler) Schedule(j *pbe.Job) *scheduler.Offer {
	return scheduler.NewOffer(m.worker, j, scheduler.Scores{})
}

func newMockSchedulerServer() *mockSchedulerServer {
	srv := server_mocks.NewMockServer()

	wconf := config.WorkerDefaultConfig()
	wconf.ServerAddress = srv.Conf.ServerAddress
	wconf.ID = "test-worker"

	// Create a worker
	x, werr := NewWorker(wconf)
	if werr != nil {
		panic(werr)
	}
	w := x.(*worker)
	// Stub the job runner so it's a no-op runner
	// i.e. ensure docker run, file copying, etc. doesn't actually happen
	w.runJob = noopRunJob

	// Create a mock scheduler with a single worker
	sched := &mockScheduler{&pbr.Worker{
		Id:    "test-worker",
		State: pbr.WorkerState_Alive,
		Jobs:  map[string]*pbr.JobWrapper{},
	}}

	m := mockSchedulerServer{srv.DB, srv, sched, srv.Conf, w}
	return &m
}

type mockSchedulerServer struct {
	db     *server.TaskBolt
	server *server_mocks.MockServer
	sched  *mockScheduler
	conf   config.Config
	worker *worker
}

func (m *mockSchedulerServer) Flush() {
	scheduler.ScheduleChunk(m.db, m.sched, m.conf)
	m.worker.checkJobs()
}

func (m *mockSchedulerServer) Close() {
	m.server.Close()
}
