let
  gitRev = "988cc958c57ce4350ec248d2d53087777f9e1949";
  pkgs = import (fetchTarball "https://github.com/NixOS/nixpkgs/archive/${gitRev}.tar.gz") { };
in
pkgs.mkShell {
  buildInputs = with pkgs; [
    go_1_20
    rsync
    docker
  ];

  shellHook = ''
    unset GOROOT
  '';
}
