{
  description = "GCP Dataflow setup";

  inputs.nixpkgs.url = "github:NixOS/nixpkgs/35f1f865c03671a4f75a6996000f03ac3dc3e472";
  inputs.flake-utils.url = "github:numtide/flake-utils";

  outputs = { self, nixpkgs, flake-utils }:
    flake-utils.lib.eachDefaultSystem (system:
      let
        pkgs = nixpkgs.legacyPackages.${system};
        beam = import ./beam.nix { inherit pkgs; };
        shell = pkgs.mkShell {
          packages = [
            pkgs.gcc-unwrapped.lib
            pkgs.poetry
            beam.beam
          ];
        };
      in
      {
        beam = beam.beam;
        # nix shell
        defaultPackage = beam.beam;
        # nix develop
        devShells.default = shell;
      });
}
