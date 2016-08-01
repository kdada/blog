(function (global) {
  System.config({
    map: {
      'app': '/app/js',
      '@angular': '/app/js/lib/angular',
      'rxjs': '/app/js/lib/rxjs',
    },
    packages: {
      'app': {defaultExtension: 'js' },
      '@angular': { defaultExtension: 'umd.min.js' },
      'rxjs': { defaultExtension: 'js' },
    },
  });
})(this);