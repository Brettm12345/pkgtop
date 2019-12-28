with import <nixpkgs> { };

buildGoPackage rec {
  name = "pkgtop";
  goPackagePath = "github.com/orhun/pkgtop";
  src = ./src;
  buildInputs = [ makeWrapper ];
  goDeps = ./deps.nix;
  meta = with stdenv.lib; {
    description =
      "Interactive package manager and resource monitor designed for the GNU/Linux.";
    license = licenses.gpl3;
    homepage = "https://github.com/orhun/pkgtop";
  };
}
