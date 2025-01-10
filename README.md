# KnockKnock 🚪
A high-performance web fuzzing tool built in Go for security testing and directory discovery.

## Overview 🔍
KnockKnock is a concurrent web fuzzing tool that helps security researchers and penetration testers discover hidden endpoints and directories on web servers. It features customizable wordlist-based fuzzing with rate limiting and parallel processing capabilities.

## Features 🌟
* Concurrent Fuzzing: Multi-threaded architecture for fast scanning 🚀
* Rate Limiting: Configurable requests per second to avoid overwhelming servers ⚡
* Custom Wordlists: Support for user-defined wordlist files 📝
* Smart Error Handling: Robust error reporting and status code filtering 🎯
* Resource Efficient: Managed concurrency with worker pools 💪

## Installation 📥
`go get github.com/Mr-b0nzai/KnockKnock`

### Usage 🛠️
Basic command syntax:
`knockknock -url "http://example.com/FUZZ" -w wordlist.txt [options]`

Parameters
* -url: Target URL with FUZZ keyword (required)
* -w: Path to wordlist file (required)
* -t: Number of concurrent workers (default: 10)
* -rate: Requests per second (default: 10)

Example:
`knockknock -url "http://example.com/FUZZ" -w directories.txt -t 20 -rate 15`

### License 📄
This project is licensed under the MIT License - see the LICENSE file for details.

### Contributing 🤝
Contributions are welcome! Please feel free to submit pull requests.

#### Security Considerations ⚠️
Please use this tool responsibly and only on systems you have permission to test.

Note: This tool is intended for legal security testing purposes only.