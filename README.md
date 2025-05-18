# Lazysh

Ever opened your terminal and thought your shell starts slow?

Well, that probably happens because of the enormous amounts of version managers, etc. which are loaded in your *rc-file.

The solution: lazy load that stuff!!1!

> [!NOTE]
> Currently only works with [fish](https://fishshell.com/). Support for other shells is wip.

## Installation

### From source

To install the `lazysh`-binary on your system, simply run:

~~~sh
sudo make clean install
~~~

## Getting started

Now, in your *rc-file (`.config/fish/config.fish`, `.zshrc`, `.bashrc`, ...) paste in the following:

~~~sh
source $(echo '
<init commands>
' | lazysh)
~~~

and replace `<init commands>` with your init commands, like `rbenv init - | source`, each placed on a new line.

It could look like:

~~~sh
source $(echo '
zoxide init fish | source
rbenv init - | source
' | lazysh)
~~~

That's it!
