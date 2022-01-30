package sanity

import (
	"fmt"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"strings"
)

// item: app's flags
type item interface {
	// label: get item's name
	label() string
	// check: whether the userSetting value is equal to the conflicts value
	check() bool
	// verbose: show the readable flag
	verbose() string
}

type boolItem struct {
	name  string
	value bool
}

func (b boolItem) label() string {
	return b.name
}

func (b boolItem) check() bool {
	return viper.GetBool(b.label()) == b.value
}

func (b boolItem) verbose() string {
	return fmt.Sprintf("--%v=%v", b.name, b.value)
}

type stringItem struct {
	name  string
	value string
}

func (s stringItem) label() string {
	return s.name
}

func (s stringItem) check() bool {
	return strings.ToLower(viper.GetString(s.label())) == s.value
}

func (s stringItem) verbose() string {
	return fmt.Sprintf("--%v=%v", s.name, s.value)
}

// conflictPair: configA and configB are conflict pair
type conflictPair struct {
	configA item
	configB item
}

// checkConflict: check configA vs configB by cmd and viper
// if both configA and configB is set by user,
// and the value is equal to the conflicts value then complain it
func (cp *conflictPair) checkConflict(cmd *cobra.Command) error {
	if checkUserSetFlag(cmd, cp.configA.label()) &&
		checkUserSetFlag(cmd, cp.configB.label()) {
		if cp.configA.check() &&
			cp.configB.check() {
			return fmt.Errorf(" %v conflict with %v", cp.configA.verbose(), cp.configB.verbose())
		}
	}

	return nil
}

// checkUserSetFlag If the user set the value (or if left to default)
func checkUserSetFlag(cmd *cobra.Command, inFlag string) bool {
	flag := cmd.Flags().Lookup(inFlag)
	return flag != nil && flag.Changed
}
