;(function() {
  // app.js
  'use strict';

  angular.module('itrak', [
    'ui.router',
    'ngAnimate',
    'ngResource',

    //foundation
    'foundation',
    'foundation.dynamicRouting',
    'foundation.dynamicRouting.animations'
  ])
    .config(config)
    .run(run)
    .constant('ServerName', 'http://localhost:8082')
    .filter('unsafe', function($sce) { return $sce.trustAsHtml; })
    ;
  
  config.$inject = ['$urlRouterProvider', '$locationProvider'];

  function config($urlProvider, $locationProvider) {
    $urlProvider.otherwise('/');

    $locationProvider.html5Mode({
      enabled:false,
      requireBase: false
    });

    $locationProvider.hashPrefix('!');
  }

  function run($rootScope) {
    FastClick.attach(document.body);

    $rootScope.$on('$stateChangeStart', function (event, toState) {
      console.log('RootScope sees',event,toState)
    })    
  }  

})();
