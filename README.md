# Urban Dictionary CLI
Search the meaning of a word or text on urban and print it on terminal :notebook:

>>> ![image](https://user-images.githubusercontent.com/71683721/175098387-02cd8f81-1c77-44b5-bcea-cc2a70e03c72.png)

## Installation
### From source
You need to compile urban cli by following the steps provided below.

* Clone the repo
```sh
# You must have to gut preinstalled
git clone https://github.com/Ruzie/urban.git
```

* Install Go and make
```sh
# For Debian (Ubuntu, Xubuntu etc)
apt install golang make

# For Arch (Endeavour, Manjaro etc)
pacman -S golang make

# For macOS
brew install golang make
```
* Compile
```sh
# Change current directory to the cloned one and run make command
cd urban && make
```
And that's it! Make should create a build folder, move there and run the program.

### Prebuild binary
Head over to [releases](https://github.com/Ruzie/urban/releases) and grab the lastest one.

Alternatively you can also do
```sh
# This is pull the lastest one from releases
wget https://github.com/Ruzie/urban/releases/latest/download/urban.xz
```
##### Note: All prebuild binary files were build against x86_64 architecture and GLIBC 2.31 or above.

## Usage
```
urban [options] [word]

Options:
 -h, shows help menu
 -d, definition of a word or text
 -e, examples of a word or text
 -j, raw json stdout
 -r, get random definitions or examples (must use as a second param)
 
Examples:
  urban -d sample (param -d refers definition(s))
  urban -e sample (param -e refers example(s))
  urban -j sample (param -j refers JSON stdout)
  urban -d -r (param -r after -d refers that you're requesting for random definition(s))
  urban -e -r (param -e after -d refers that you're requesting for random example(s))
  urban -j -r (param -j after -d refers that you're requesting for random JSON body of the result)
```