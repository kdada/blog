(function (global) {
  System.config({
    map: {
      'app': '/app/js',
      '@angular': '/app/js/lib/angular'
    },
    packages: {
      'app': {defaultExtension: 'js' },
      '@angular': { defaultExtension: 'umd.min.js' }
    },
  });
})(this);