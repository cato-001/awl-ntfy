{ pkgs }:

pkgs.mkShell {
  packages = [
    pkgs.mkGoEnv { pwd = ./.; }
    pkgs.gomod2nix
  ];
}
