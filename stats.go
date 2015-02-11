package docker

import "io"

type StatsOptions struct {
	Id           string
	OutputStream io.Writer
	ErrorStream  io.Writer
	RawTerminal  bool
	Success      chan struct{}
}

func (c *Client) Stats(opts StatsOptions) error {

	if opts.Id == "" {
		return &NoSuchContainer{ID: opts.Id}
	}

	path := "/containers/" + opts.Id + "/stats"

	//return c.stream(method, path, setRawTerminal, rawJSONStream, headers, in, stdout, stderr)
	//return c.stream("GET", path, opts.RawTerminal, opts.RawJSONStream, nil, nil, opts.OutputStream, nil)

	//return c.hijack(method, path, success, setRawTerminal, in, stderr, stdout, data)
	return c.hijack("GET", path, opts.Success, opts.RawTerminal, nil, opts.ErrorStream, opts.OutputStream, nil)
}

/*

type APIStats struct {
	Read    time.Duration `json:"read"`
	Network APINetworkStats
	Memory  APIMemoryStats
	Blkio   APIBlkioStats
	Cpu     APICpuStats
}

type APICpuStats struct{}

type APIBlkioStats struct{}

type APIMemoryStats struct{}

type APINetworkStats struct{}

*/
