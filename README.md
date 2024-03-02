# Password Generator

A simple password generator written in Go.

As Posted on [ebourgess.dev](https://ebourgess.dev/posts/creating-a-password-generator/)

## Description

This script generates random passwords based on user-defined criteria such as length and character types.

## Features

- Customizable password length
- Option to include uppercase letters, lowercase letters, numbers, and special characters
- Secure and random password generation

## Installation

1. Clone the repository:

    ```shell
    git clone https://github.com/ebourgess/pass-generator.git
    ```

2. Build the executable:

    ```shell
    go build pass-generator.go
    ```

3. To use it globally move it to `/usr/local/bin`

    ```shell
    sudo mv pass-generator /usr/local/bin
    ```

## Usage

```shell
pass-generator -h
    --digits
        Use digits
    --length int
        Password length (default 12)
    --lower
        Use lowercase letters (default true)
    --special
        Use special characters
    --upper
        Use uppercase letters (default true)
```

In order to use this simply run the command

```shell
pass-generator --length 32 --digits --lower --special --upper
```

Keep in mind that by default, `length` is 12, `lower` and `upper` are both enabled. So if you need to add `digits` or `special` characters, you will need to add the flags `special` and `upper` to your command.

## Test your new password

In order to test your password's strength, copy and paste it [here](https://www.security.org/how-secure-is-my-password/)

## License

This project is licensed under the MIT License. See the [LICENSE](LICENSE) file for more information.
