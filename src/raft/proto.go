package raft

import "labrpc"
import "sync"

//
// as each Raft peer becomes aware that successive log entries are
// committed, the peer should send an ApplyMsg to the service (or
// tester) on the same server, via the applyCh passed to Make(). set
// CommandValid to true to indicate that the ApplyMsg contains a newly
// committed log entry.
//
// in Lab 3 you'll want to send other kinds of messages (e.g.,
// snapshots) on the applyCh; at that point you can add fields to
// ApplyMsg, but set CommandValid to false for these other uses.
//
type ApplyMsg struct {
	CommandValid bool
	Command      interface{}
	CommandIndex int
}

type LogEntry struct {
  Term int
	Cmd interface{}
}

// Used to wake up a goroutine
// false: wait up and exit
// true: wait up and continue work
type WakeupChan chan bool

// A Go object implementing a single Raft peer.
//
type Raft struct {
  // Ready-only members
	peers     []*labrpc.ClientEnd // RPC end points of all peers
	persister *Persister          // Object to hold this peer's persisted state
	me        int                 // this peer's index into peers[]
  applyCh chan ApplyMsg

  sync.Mutex
  // For all servers
  currentTerm int
  votedFor int
  // The first entry is a sentinel
  log []LogEntry
  leader int

  commitIndex int
  lastApplied int

  // For the leader
  nextIndex []int
  matchIndex []int

  // --- The members above are protectd by the 'Mutex'
  af *AsyncFSA

  applierWakeup WakeupChan

  killed bool
}

// example RequestVote RPC arguments structure.
// field names must start with capital letters!
//
type RequestVoteArgs struct {
	// Your data here (2A, 2B).
  Term int
  CandidateId int
  LastLogIndex int
  LastLogTerm int
}

// For both types of RPCs
type RequestReply struct {
  // currentTerm, for leader to update itself
  Term int
  // AppendEntries(): true if follower contained entry matching prevLogIndex
  //    and prevLogTerm
  // RequestVote(): true if success
  Success bool
  // Which server is this reply from?
  Peer int
  // How many entries of the leader at the moment the request was sent?
  NumLeaderLogEntries int
}

type AppendEntriesArgs struct {
  Term int // leader's term
  LeaderId int // so follower can redirect clients
  PrevLogIndex int // index of log entry immediately preceding new ones
  PrevLogTerm int // term of prevLogIndex entry
  // log entries to store (empty for heartbeat; may send more than one for
  // efficiency)
  Entries []LogEntry
  LeaderCommit int // leader's commitIndex
}


