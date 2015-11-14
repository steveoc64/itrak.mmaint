;(function() {
  // app.js
  'use strict';

  angular.module('itrak', [
    'ui.router',
    'ngAnimate',
    'ngResource',

    //foundation
    'foundation',
    //'foundation.dynamicRouting',
    //'foundation.dynamicRouting.animations'
  ])
    .config(config)
    .run(run)
    .constant('ServerName', '')
    .filter('unsafe', function($sce) { return $sce.trustAsHtml; })
    ;
  
  config.$inject = ['$stateProvider','$urlRouterProvider', '$locationProvider'];

  function config($stateProvider, $urlRouterProvider, $locationProvider) {
    $urlRouterProvider.otherwise('/');

    $locationProvider.html5Mode({
      enabled:false,
      requireBase: false
    });

    $locationProvider.hashPrefix('!');

    // Manually create all the routes here
    $stateProvider
      .state('login',{
        url: '/',
        templateUrl: 'templates/login.html',
        controller: 'loginCtrl',
        controllerAs: 'loginCtrl',
        animation: {
          enter: 'hingeInFromTop'
        },
      })
        .state('admin',{
          url: '/admin',
          templateUrl: 'templates/admin.html',
          controller: 'adminCtrl',
          controllerAs: 'adminCtrl',
          animation: {
            enter: 'slideInRight'
          }
        })
          .state('admin.equipment',{
            url: '/equipment',
            templateUrl: 'templates/admin.equipment.html',
            controller: 'adminEquipmentCtrl',
            controllerAs: 'adminEquipmentCtrl',
            resolve: {
              equipments: function(Equipment) {
                return Equipment.query()
              },
              sites: function(Sites) {
                return Sites.query()
              },
              vendors: function(Vendors) {
                return Vendors.query()
              }
            }
          })
          .state('admin.equipment_details',{
            url: '/equipdetails/:id',
            templateUrl: 'templates/admin.equipment_details.html',
            controller: 'adminEquipmentDetCtrl',
            controllerAs: 'adminEquipmentDetCtrl',
            resolve: {
              equipment: function($stateParams,Equipment) {
                return Equipment.get({id: $stateParams.id})
              },
              subparts: function($stateParams,SubParts) {
                return SubParts.query({id: $stateParams.id})
              },
              sites: function(Sites) {
                return Sites.query()
              },
              vendors: function(Vendors) {
                return Vendors.query()
              }
            }
          })
  }

  function run($rootScope) {
    FastClick.attach(document.body);
  }  

})();
