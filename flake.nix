{
  description = "Development flake for Gize";

  inputs = {
    nixpkgs.url = "github:nixos/nixpkgs/nixos-24.11";
  };

  outputs =
    { nixpkgs, ... }:
    let
      x86 = "x86_64-linux";
      pkgs = nixpkgs.legacyPackages."${x86}";
    in
    {
      devShells."${x86}".default = pkgs.mkShell {
        packages = with pkgs; [
          # Golang
          go
          gotools

          # Formatters
          treefmt
          beautysh
          mdformat
          yamlfmt
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
      };
    };
}
