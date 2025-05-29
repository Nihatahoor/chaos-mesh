// +build s390x

package ptrace

func getRegs() error {
    return fmt.Errorf("ptrace not supported on s390x")
}

