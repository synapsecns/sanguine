package dbcommon

import (
	"fmt"
	"strings"
)

// MysqlRealEscapeString mimics the behavior of the PHP function
// mysql_real_escape_string, which is used to escape special characters in a
// string before sending it to a MySQL database. This can help prevent SQL
// injection attacks.
//
// The PHP function itself has a checkered history. While its primary intention
// was to make strings safe for MySQL, relying solely on it is considered bad
// practice. This is because the best way to prevent SQL injection is by using
// prepared statements or parameterized queries, which don't require manual
// string escaping. Additionally, mysql_real_escape_string in PHP relies on the
// current character set, and without the proper character set, it can fail to
// protect against SQL injection.
//
// It's worth noting that using this Go implementation is, in a way, an
// acknowledgement of an anti-pattern from PHP. While it can escape some
// characters and may prevent some naive SQL injection attempts, it is not a
// substitute for  prepared statements or parameterized queries.
// This implementation is provided as a convenience for some legacy code and
// should not be considered a robust security solution. Furthermore, it is only used on blockchain
// data which is all public. Do not use this in any situation where you are dealing w/ private data.
//
// Deprecated: Use prepared statements or parameterized queries instead going forward.
func MysqlRealEscapeString(value string) string {
	var sb strings.Builder
	for i := 0; i < len(value); i++ {
		c := value[i]
		switch c {
		case '\\', 0, '\n', '\r', '\'', '"':
			sb.WriteByte('\\')
			sb.WriteByte(c)
		case '\032':
			sb.WriteByte('\\')
			sb.WriteByte('Z')
		default:
			sb.WriteByte(c)
		}
	}
	return sb.String()
}

// SprintfEscape is a custom implementation of the fmt.SprintfEscape function,
// with a twist. While it allows the formatting of a wide range of
// data types into a string, it takes special care when dealing with
// string arguments.
//
// Any string argument passed to this function is first passed through
// the MysqlRealEscapeString function to escape special characters that
// could be problematic or even dangerous in the context of MySQL
// queries. This behavior mimics the PHP function mysql_real_escape_string,
// which was historically used to prevent SQL injection attacks. While
// this might provide some protection against naive SQL injection attempts,
// it is essential to note that the best security practice for databases
// is to use prepared statements or parameterized queries, rather than
// manually escaping strings.
//
// This function is particularly useful when working with legacy systems
// or in situations where prepared statements might not be feasible.
// However, developers should always be aware of its limitations and
// consider it a convenience, not a robust security solution.
//
// Deprecated: Use prepared statements or parameterized queries instead going forward.
func SprintfEscape(format string, a ...interface{}) string {
	escapedArgs := make([]interface{}, len(a))
	for i, arg := range a {
		switch v := arg.(type) {
		case string:
			escapedArgs[i] = MysqlRealEscapeString(v)
		case []string:
			escapedStrings := make([]string, len(v))
			for it, s := range v {
				escapedStrings[it] = MysqlRealEscapeString(s)
			}
			escapedArgs[i] = escapedStrings
		default:
			escapedArgs[i] = v
		}
	}
	return fmt.Sprintf(format, escapedArgs...)
}
