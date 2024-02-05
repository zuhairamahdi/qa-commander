'use strict';

const EmberApp = require('ember-cli/lib/broccoli/ember-app');

module.exports = function (defaults) {
  const app = new EmberApp(defaults, {
    // Add options here
  });
  app.import('public/libs/jquery-3.7.1.min.js')
  app.import('public/js/app.js')
  return app.toTree();
};
