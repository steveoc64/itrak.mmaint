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
          .state('admin.people',{
            url: '/people',
            templateUrl: 'templates/admin.people.html',
            controller: 'adminPeopleCtrl',
            controllerAs: 'adminPeopleCtrl',
            resolve: {
              people: function(People) {
                return People.query()
              },
              sites: function(Sites) {
                return Sites.query()
              },
              roles: function(Roles) {
                return Roles.query()
              }
            }
          })
          .state('admin.site',{
            url: '/site',
            templateUrl: 'templates/admin.site.html',
            controller: 'adminSiteCtrl',
            controllerAs: 'adminSiteCtrl',
            resolve: {
              sites: function(Sites) {
                return Sites.query()
              },
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
          .state('admin.spares',{
            url: '/spares',
            templateUrl: 'templates/admin.spares.html',
            controller: 'adminSparesCtrl',
            controllerAs: 'adminSparesCtrl',
            resolve: {
              spares: function(Spares) {
                return Spares.query()
              },
              vendors: function(Vendors) {
                return Vendors.query()
              },
              sites: function(Sites) {
                return Sites.query()
              }

            }
          })
          .state('admin.consumables',{
            url: '/cons',
            templateUrl: 'templates/admin.consumables.html',
            controller: 'adminConsumablesCtrl',
            controllerAs: 'adminConsumablesCtrl',
            resolve: {
              consumables: function(Consumables) {
                return Consumables.query()
              },
              vendors: function(Vendors) {
                return Vendors.query()
              },
              sites: function(Sites) {
                return Sites.query()
              }

            }
          })
          .state('admin.equiptypes',{
            url: '/equiptypes',
            templateUrl: 'templates/admin.equiptypes.html',
            controller: 'adminEquipTypesCtrl',
            controllerAs: 'adminEquipTypesCtrl',
            resolve: {
              equiptypes: function(EquipTypes) {
                return EquipTypes.query()
              },
              vendors: function(Vendors) {
                return Vendors.query()
              },
              sites: function(Sites) {
                return Sites.query()
              }

            }
          })
          .state('admin.task',{
            url: '/task',
            templateUrl: 'templates/admin.task.html',
            controller: 'adminTaskCtrl',
            controllerAs: 'adminTaskCtrl',
            resolve: {
              tasks: function(Tasks) {
                return Tasks.query()
              }
            }
          })
          .state('admin.sitetask',{
            url: '/sitetask/:id',
            templateUrl: 'templates/admin.sitetask.html',
            controller: 'adminSiteTaskCtrl',
            controllerAs: 'adminSiteTaskCtrl',
            resolve: {
              site: function($stateParams,Sites) {
                return Sites.get({id: $stateParams.id})
              },
              tasks: function($stateParams,SiteTasks) {
                return SiteTasks.query({id: $stateParams.id})
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
          .state('admin.vendor',{
            url: '/vendor/:id',
            templateUrl: 'templates/admin.vendor.html',
            controller: 'adminVendorCtrl',
            controllerAs: 'adminVendorCtrl',
            resolve: {
              vendors: function(Vendors) {
                return Vendors.query()
              }
            }
          })
          .state('admin.workorder',{
            url: '/workorder',
            controller: 'workerCtrl',
            controllerAs: 'workerCtrl',
            template: '<h3>Work Orders</h> ... tbd'
          })
        .state('worker',{
          url: '/worker',
          templateUrl: 'templates/worker.html',
          controller: 'workerCtrl',
          controllerAs: 'workerCtrl',
          animation: {
            enter: 'slideInRight'
          }
        })
          .state('worker.dashboard',{
            url: '/dashboard',
            template: '<h3>Worker Dashboard</h> ... tbd'
          })
          .state('worker.estop',{
            url: '/estop',
            templateUrl: 'templates/worker.estop.html',
            controller: 'workerStopCtrl',
            controllerAs: 'workerStopCtrl',
            resolve: {
              siteEquip: function(SiteEquipment, loginState) {
                return SiteEquipment.query({id: loginState.site})
              }
            }
          })
          .state('worker.pstop',{
            url: '/pstop',
            templateUrl: 'templates/worker.pstop.html',
            controller: 'workerPStopCtrl',
            controllerAs: 'workerPStopCtrl',
            resolve: {
              siteEquip: function(SiteEquipment, loginState) {
                return SiteEquipment.query({id: loginState.site})
              }
            }
          })
          .state('worker.equip',{
            url: '/equip',
            templateUrl: 'templates/worker.equip.html',
            controller: 'workerEquipmentCtrl',
            controllerAs: 'workerEquipmentCtrl',
            resolve: {
              siteEquip: function(SiteEquipment, loginState) {
                return SiteEquipment.query({id: loginState.site})
              }
            }
          })
          .state('worker.task',{
            url: '/task',
            template: '<h3>Uncompleted Tasks for this Site</h> ... tbd'
          })
          .state('worker.spares',{
            url: '/spares',
            template: '<h3>Spare Parts at this Site</h> ... tbd'
          })
          .state('worker.reports',{
            url: '/reports',
            template: '<h3>Maintenance Reports</h> ... tbd'
          })
  }

  function run($rootScope) {
    FastClick.attach(document.body);
  }  

})();
