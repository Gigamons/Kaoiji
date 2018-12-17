package consts

type LoginReply int32

const (
	LoginFailed LoginReply = -iota-1    // -1
	LoginClientOutdated                 // -2
	LoginBanned                         // -3
	LoginMultiAcc                       // -4
	LoginException                      // -5
	LoginRequireSupporter               // -6
	LoginPasswordReset                  // -7
	LoginRequireVerification            // -8
)
