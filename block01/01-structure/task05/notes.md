1. -rwxrwxr-x 1 peekaboo 2.3M Sep 10 14:58 app
2. It didn't output anything.
3. It didn't output anything.
4. func Println(a ...any) (n int, err error)
   Println formats using the default formats for its operands and writes to
   standard output. Spaces are always added between operands and a newline
   is appended. It returns the number of bytes written and any write error
   encountered.
5. 2.3 MB → 1.5 MB after `-s -w`; this means symbols and DWARF account for 0.8 MB—about a third.
   The remaining 1.5 MB cannot be removed—it consists of the Go runtime and the parts of the standard library that are actually used.
   The binary is statically linked and has no external dependencies—hence both its size and its portability.
6. -rwxrwxr-x 1 peekaboo 2.3M Sep 10 15:38 bin/app
   -rwxrwxr-x 1 peekaboo 1.5M Sep 10 15:38 bin/app-small
