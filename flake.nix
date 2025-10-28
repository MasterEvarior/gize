{
  description = "Development flake for Gize";

  inputs = {
    nixpkgs.url = "github:nixos/nixpkgs/nixos-25.05";
  };

  outputs =
    { self, nixpkgs, ... }:
    let
      x86 = "x86_64-linux";
      pkgs = nixpkgs.legacyPackages."${x86}";
      lib = pkgs.lib;
      goPackages = with pkgs; [
        go
        golangci-lint
      ];
      formatterPackages = with pkgs; [
        treefmt
        beautysh
        mdformat
        yamlfmt
        jsonfmt
        deadnix
        nixfmt-rfc-style
      ];
      vendorHash = "sha256-/OzNsgU3VNnkL9sXDoZahJ7fMqoYCEmstnNnGvmF03A=";
    in
    {
      devShells."${x86}".default = pkgs.mkShellNoCC {
        packages = goPackages ++ formatterPackages;

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

      packages."${x86}" = {
        default = pkgs.buildGoModule {
          inherit vendorHash;

          pname = "gize";
          version = "v1.1.0";
          src = ./.;

          meta.mainProgram = "gize";
        };

        image = pkgs.dockerTools.buildLayeredImage {
          name = "gize";

          contents = [ self.packages.${x86}.default ];

          config.Cmd = [ "${lib.getExe self.packages.${x86}.default}" ];
        };
      };

      apps.${x86} = {
        default = {
          type = "app";
          program = "${lib.getExe self.packages.x86_64-linux.default}";
          meta.description = "Run Gize, a tool to display Git repositories through a website.";
        };

        image = {
          type = "app";
          program = lib.getExe (
            pkgs.writeShellApplication {
              name = "gize";
              runtimeInputs = [ pkgs.docker ];
              text = ''
                set -euo pipefail
                image_path="${self.packages.${x86}.image}"

                echo "--- Loading Docker image: gize (from $image_path) ---"
                docker load < "$image_path"

                echo "--- Running container 'gize' on http://localhost:8080 ---"

                docker run \
                  -e GIZE_ROOT="/repositories" \
                  -v ./..:/repositories \
                  -p 8080:8080 \
                  --rm \
                  gize:${toString self.packages.${x86}.image.imageTag}
              '';
            }
          );
          meta.description = "Run Gize, a tool to display Git repositories through a website (but in a container).";
        };
      };

      checks."${x86}" = {
        tests = pkgs.buildGoModule {
          inherit vendorHash;

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
        };

        formatting = pkgs.buildGoModule {
          inherit vendorHash;

          name = "formatting";
          src = ./.;
          doCheck = true;

          nativeBuildInputs = goPackages ++ formatterPackages;

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
        };
      };
    };
}
