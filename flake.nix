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
          echo "Started Gize dev shell"
        '';
      };
    };
}
