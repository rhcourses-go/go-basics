# include build settings and scripts
do './buildscripts.pl';

############################################
# Override the default settings if needed. #
# (cf. buildscripts.pl)                    #
############################################
# $srcdir = "$rootdir/texsrc";
# $examplesdir = "$srcdir/examples";
# $packagesdir  = "$srcdir/packages";
# $builddir = "$rootdir/build"; # where to put the pdf files created.
# $buildoptions = "-interaction=nonstopmode";

###################################
# Add files to the build process. #
###################################

addfilesfromdir('texsrc/slides',
    'go-intro-handout.tex',
    'go-intro-presentation.tex',
);
