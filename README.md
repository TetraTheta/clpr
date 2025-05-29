# clpr

Utility for managing string data of clipboard

If the system does not have clipboard feature, it will use 'Local Clipboard Cache', which uses local text files as 'clipboard' instead.

## Usage

```plaintext
>clpr -h
clpr - Universal Clipboard Utility

  Usage:
    clpr [get|list|set]

  Subcommands:
    get    Get content of clipboard or LCC
    list   List stored LCC
    set    Set content of clipboard or LCC

  Flags:
       --version   Displays the program version string.
    -h --help      Displays help with available flag, subcommand, and positional value parameters.
```
```plaintext
>clpr get -h
get - Get content of clipboard or LCC

  Flags:
       --version   Displays the program version string.
    -h --help      Displays help with available flag, subcommand, and positional value parameters.
    -name --n         Name of LCC
```
```plaintext
>clpr list -h
list - List stored LCC

  Flags:
       --version   Displays the program version string.
    -h --help      Displays help with available flag, subcommand, and positional value parameters.
```
```plaintext
>clpr set -h
set - Set content of clipboard or LCC

  Usage:
    set [content]

  Positional Variables:
    content   New content of clipboard or LCC. If not present, piped string value will be used.
  Flags:
       --version   Displays the program version string.
    -h --help      Displays help with available flag, subcommand, and positional value parameters.
    -name --n         Name of LCC
```
