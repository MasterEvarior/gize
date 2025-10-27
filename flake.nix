{
  description = "Development flake for Gize";

  inputs = {
    nixpkgs.url = "github:nixos/nixpkgs/nixos-25.05";
  };

  outputs =
    { nixpkgs, ... }:
    let
      x86 = "x86_64-linux";
      pkgs = nixpkgs.legacyPackages."${x86}";
      lib = pkgs.lib;
    in
    {
      devShells."${x86}".default = pkgs.mkShellNoCC {
        packages = with pkgs; [
          # Golang
          go
          golangci-lint

          # Formatters
          treefmt
          beautysh
          mdformat
          yamlfmt
          jsonfmt
          deadnix
          nixfmt-rfc-style
        ];

        shellHook = ''
          git config --local core.hooksPath .githooks/
        '';

        # Environment Variables
        GIZE_ROOT = "./..";
        GIZE_TITLE = "Gize (dev)";
        GIZE_DESCRIPTION = "You local Git repository browser (dev)";
        GIZE_FOOTER = "Made with ❤️ and published on <a href='https://github.com/MasterEvarior/gize'>GitHub</a> (dev)";
        GIZE_PORT = ":8080";
        GIZE_ENABLE_DOWNLOAD = "true";
        GIZE_ENABLE_CACHE = "true";
      };

      packages."${x86}".default = pkgs.buildGoModule {
        pname = "gize";
        version = "v1.1.0";
        src = ./.;
        vendorHash = "sha256-/OzNsgU3VNnkL9sXDoZahJ7fMqoYCEmstnNnGvmF03A=";
      };

      checks."${x86}" = {
        tests = pkgs.buildGoModule {
          name = "tests";
          src = ./.;
          doCheck = true;

          nativeBuildInputs = [
            pkgs.go
          ];

          checkPhase = ''
            ${lib.getExe pkgs.go} test ./...
          '';

          installPhase = ''
            mkdir "$out"
          '';

          vendorHash = "sha256-/OzNsgU3VNnkL9sXDoZahJ7fMqoYCEmstnNnGvmF03A=";
        };

        formatting = pkgs.buildGoModule {
          name = "formatting";
          src = ./.;
          doCheck = true;

          nativeBuildInputs = with pkgs; [
            go
            golangci-lint

            treefmt
            beautysh
            mdformat
            yamlfmt
            jsonfmt
            deadnix
            nodePackages.prettier
            nixfmt-rfc-style
          ];

          checkPhase = ''
            export HOME=$PWD
            export GOCACHE=$PWD/.go-cache
            mkdir -p $GOCACHE

            ${lib.getExe pkgs.golangci-lint} run
            ${lib.getExe pkgs.treefmt} --ci
          '';

          installPhase = ''
            mkdir "$out"
          '';

          vendorHash = "sha256-/OzNsgU3VNnkL9sXDoZahJ7fMqoYCEmstnNnGvmF03A=";
        };
      };
    };
}
