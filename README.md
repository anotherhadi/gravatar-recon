# Gravatar-Recon 👤

<p>
    <a href="https://github.com/anotherhadi/gravatar-recon/releases"><img src="https://img.shields.io/github/release/anotherhadi/gravatar-recon.svg" alt="Latest Release"></a>
    <a href="https://pkg.go.dev/github.com/anotherhadi/gravatar-recon?tab=doc"><img src="https://godoc.org/github.com/anotherhadi/gravatar-recon?status.svg" alt="GoDoc"></a>
    <a href="https://goreportcard.com/report/github.com/anotherhadi/gravatar-recon"><img src="https://goreportcard.com/badge/github.com/anotherhadi/gravatar-recon" alt="GoReportCard"></a>
</p>

- [🧾 Project Overview](#-project-overview)
- [🚀 Features](#-features)
- [⚠️ Disclaimer](#%EF%B8%8F-disclaimer)
- [📦 Installation](#-installation)
- [🧪 Usage](#-usage)
- [💡 Examples](#-examples)
- [🤝 Contributing](#-contributing)

## 🧾 Project Overview

Retrieve and aggregate public **OSINT data from Gravatar**. Given an **email
address**, the tool queries the Gravatar API and extracts useful information
such as profile metadata, avatar, social accounts, and contact info.

## 🚀 Features

- Export results to JSON
- Fetch avatar directly
- Extract profile metadata:
  - Display name, bio, location, job, company
  - Public emails
  - Phone numbers
  - Contact forms, calendars
- Enumerate linked accounts (Twitter, GitHub, Mastodon, etc.)
- Parse profile background colors and styles

## ⚠️ Disclaimer

This tool is intended for **educational purposes only**. Use responsibly and
ensure you have permission to access the data you are querying.

## 📦 Installation

### With Go

```bash
go install github.com/anotherhadi/gravatar-recon@latest
```

### With Nix/NixOS

<details>
<summary>Click to expand</summary>

**From anywhere (using the repo URL):**

```bash
nix run github:anotherhadi/gravatar-recon -- [--flags value] target_email
```

**Permanent Installation:**

```bash
# add the flake to your flake.nix
{
  inputs = {
    gravatar-recon.url = "github:anotherhadi/gravatar-recon";
  };
}

# then add it to your packages
environment.systemPackages = with pkgs; [ # or home.packages
  inputs.gravatar-recon.defaultPackage.${pkgs.system}
];
```

</details>

## 🧪 Usage

```bash
gravatar-recon [--flags value] target_email@example.com
```

### Flags

```txt
-j, --json string     Write results to specified JSON file
-s, --silent          Suppress all non-essential output
-a, --print-avatar    Show the avatar in the output
```

## 💡 Examples

```bash
gravatar-recon myemail@gmail.com
gravatar-recon myemail@gmail.com --json output.json
```

## 🤝 Contributing

Contributions are welcome! See [CONTRIBUTING.md](CONTRIBUTING.md) for details.
