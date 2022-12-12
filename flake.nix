{
  description = "A basic flake with a shell";
  inputs.nixpkgs.url = "github:NixOS/nixpkgs/nixpkgs-unstable";
  inputs.flake-utils.url = "github:numtide/flake-utils";

  outputs = {
    self,
    nixpkgs,
    flake-utils,
  }:
    flake-utils.lib.eachDefaultSystem (system: let
      pkgs = nixpkgs.legacyPackages.${system};
    in {
      devShells.default = pkgs.mkShell {
        nativeBuildInputs = [pkgs.bashInteractive];
        buildInputs = with pkgs; [
          go_1_19
          gofumpt
          gotools
          (pkgs.buildGoModule rec {
            pname = "golines";
            version = "0.11.0";

            src = pkgs.fetchFromGitHub {
              owner = "segmentio";
              repo = "golines";
              rev = "v${version}";
              sha256 = "sha256-2K9KAg8iSubiTbujyFGN3yggrL+EDyeUCs9OOta/19A=";
            };

            vendorSha256 = "sha256-rxYuzn4ezAxaeDhxd8qdOzt+CKYIh03A9zKNdzILq18=";

            nativeBuildInputs = [pkgs.installShellFiles];
          })
        ];
      };
    });
}
