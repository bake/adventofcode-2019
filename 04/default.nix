{ pkgs ? import <nixpkgs> {} }:
pkgs.stdenv.mkDerivation {
  name = "aoc2019day04";
  src = ./.;
  buildInputs = [ pkgs.ghc ];
  buildPhase = "ghc main.hs";
  installPhase = "mkdir -p $out && cp main $out";
}
