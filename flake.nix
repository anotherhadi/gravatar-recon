{
  description = "Retrieve and aggregate public OSINT data from Gravatar. Given an email address, the tool queries the Gravatar API and extracts useful information such as profile metadata, avatar, social accounts, and contact info.";

  inputs = {nixpkgs.url = "github:NixOS/nixpkgs/nixos-unstable";};

  outputs = {
    self,
    nixpkgs,
  }: let
    supportedSystems = ["x86_64-linux" "aarch64-linux"];

    forAllSystems = f:
      nixpkgs.lib.genAttrs supportedSystems
      (system: f system (import nixpkgs {inherit system;}));

    pname = "gravatar-recon";
    version = "1.0.0";

    ldflags = ["-s" "-w"];
  in {
    packages = forAllSystems (system: pkgs: {
      "${pname}" = pkgs.buildGoModule {
        inherit pname version ldflags;

        src = ./.;
        subPackages = ["cmd"];
        outputs = ["out"];
        installPhase = ''
          mkdir -p $out/bin
          cp $GOPATH/bin/cmd $out/bin/gravatar-recon
        '';

        vendorHash = "sha256-hjaIXZMK9+b+tlWD55OU3mS0CLUA/Oonn/RBHQdgs2g=";

        meta = with pkgs.lib; {
          description = "Retrieve and aggregate public OSINT data from Gravatar. Given an email address, the tool queries the Gravatar API and extracts useful information such as profile metadata, avatar, social accounts, and contact info.";
          homepage = "https://github.com/anotherhadi/gravatar-recon";
          platforms = platforms.unix;
        };
      };
    });

    defaultPackage =
      forAllSystems (system: pkgs: self.packages.${system}.${pname});
  };
}
