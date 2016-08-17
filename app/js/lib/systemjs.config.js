(function (global) {
  System.config({
    map: {
      'app': '/js',
      '@angular': '/js/lib/angular',
    },
    packages: {
      'app': {defaultExtension: 'js' },
      '@angular': { defaultExtension: 'umd.min.js' },
    },
  });
})(this);