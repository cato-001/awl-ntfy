{ perSystem }:

perSystem.gomod2nix.buildGoApplication {
  pname = "awl-ntfy";
  version = "0.1";
  pwd = ./.;
  src = ./.;
  modules = ./gomod2nix.toml;
}
