1. -rwxrwxr-x 1 peekaboo 2.3M Sep 10 14:58 app
2. It didn't output anything.
3. It didn't output anything.
4. func Println(a ...any) (n int, err error)
   Println formats using the default formats for its operands and writes to
   standard output. Spaces are always added between operands and a newline
   is appended. It returns the number of bytes written and any write error
   encountered.
5. The large binary size for a simple program is due to the fact that the compiler includes bulky standard library code and associated debugging information by default.