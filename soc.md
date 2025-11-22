# 12.11.25

I'd like to build a Password Manager, though I hate the idea of a 'Password Manager' -- the word
'Manager' makes it seem like it's a person constantly doing shit. What I really want is a nice
little program that can store my passwords safely & securely.

Since I'm an archlinux user, it will be built with the assumption of an archlinux system.
Since I want to write go, it will be written in go.

## Interfaces

Say I am working on a project for a new company *acme*; *acme* have a lot of bullshit systems
which I need to sign up for, each one requiring a password.
How can I create & then store these passwords on my computer to enable ease of retrieval?

### Creating & Using

I would like to have a simple terminal interface -- nothing overfly flashy; something which
gets the job done. I'm going to assume the users of this, primarily me, will not be retards
requiring handholding.
I don't want too many prompts, I'd rather have optional flags.

```sh
# Creates a new password of xxx associated with jira
pword -n jira xxx

# Creates a random password and associates it with jira
pword -n jira --random

# Copies the password assoicated with jira to the clipboard
pword -c jira

# Prints the password associated with jira to standard output
pword jira
```

### Namespaces

Regarding my original problem, I stated that I was doing a contract for *acme*; what if *acme* requires
that I sign up to systems of which I am currently signed up for, wether on my personal or through
another contract?

e.g
I personally have a password for `gcloud` and acme also require a google cloud account with an
associated password.

Do I really want to change the keyname of my password for google cloud to avoid conflicts?
In my opinion, the following looks ugly:
```sh
# Prints the existing gcloud password
pword gcloud

# ERROR
pword -n gcloud -r

# Ugly
pword -n acme-gcloud -r
```

To resolve this ugliness, we can utilise a namespace idea. Where passwords can be grouped together
under a single namespace:
```sh
# Creates a new namespace for acme
pword -ns acme --create

# Lists all password keynames associated with acme
pword -ns acme --list

# Creates a new password for gcloud associated with the acme namespace
pword -ns acme -n gcloud -r
```

### Storage

Back to the problem of storing passwords...

Working with the assumption of an archlinux directory structure & with the fact that we want the
passwords created to be related to a given user, it makes the most sense to store the passwords
within the home directory of the user.

Specifically, we can store the passwords within the `.local/` directory of a given user:

```
  .local/
      pword/
          acme
          user
          wiktorg
```

Where `acme` & `wiktorg` are both namespaces, `user` is the standard file for storing passwords.

Each file can simply be a key-value mapping:

```acme
jira xytgeugb_encrypted_password_debhjkfbsjh
gcloud gsefeb_encrypted_password_fsdjknfsk
```

## Implementing

**Types**

# 22.11.25

I have just deleted the main file which I had begun as the main event loop of the pword program,
and while I had made some sense of *progress*, I had felt the cloud of complexity weining over
my mental state. It is this which contributed to my not attending to this program for 10 days now.

Complexity came, in this instance, in the attempt of implementing the command line interface which
I designed, mapping the commands to actual functionality.
In attempting to implement, I did not reason too much with regards to how I *should* implement
a cli, or how one goes about building a cli in the first place; I thought that -- being a small program
-- I could hold all of the different commands in my head.
How wrong I was...

I would now like to *think* about the procedure behind taking input in this program, and mapping
that input to functionality.
To do this, I will define the set of all *symbols* which are represented as input within my program:

```go
command = "pword" password | "pword" namespace {password} | "pword" namespace "--list"

password = "-p" word {pass_option} | "-p" word "--new" (word | "--random")

namespace = "-ns" ns {ns_option}

ns_option = "--create"
pass_option = "--copy"

word = char{char}
char = "A"..."Z" | "a"..."z" | "0"..."9"
```

We are going to assume that our language has an `os.Args` global variable, which produces
an array of whitespace-separated input values starting index after the first value of "pword".

This then enables us to iterate through `os.Args`, scanning each element of input to
it's corresponding symbol value.

From the definition above, we can infer that `os.Args` should produce an array where each
value are one of the constants defined below.

```go
type Symbol uint8

const (
    passInitFlag Symbol = iota
    word
    newFlag
    createFlag
    randomFlag
    nsInitFlag
    copyFlag
    listFlag
)
```

Scanning can then be done through the calling of `Scan()` defined below:

```go
func DetermineSymbol(symbol string) Symbol {
   switch symbol {
    case "-p":
        return passInitFlag
    case "-n":
        return newFlag
    case "--create":
        return createFlag
    case "--random":
        return randomFlag
    case "-ns":
        return nsInitFlag
    case "--list":
        return listFlag
    case "--copy":
        return copyFlag
    default:
        return word
} 
}

func Scan() []Symbol {
    var symbols []Symbol
    for i:= range os.Args {
        s := DetermineSymbol(os.Args[i])
        symbols = append(symbols, s)
    }
    return symbols
}
```
