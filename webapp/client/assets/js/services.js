;(function(){
  // Define all services in here
  'use strict';

  console.log('Loading Services')

  // People
  angular.module('itrak').factory('People', function($resource, ServerName){
    return $resource(ServerName+'/people/:id', {id:'@personId'});
  })

  // Sites
  angular.module('itrak').factory('Sites', function($resource, ServerName){
    return $resource(ServerName+'/site/:id', {id: '@id'});
  })

  // Equipment
  angular.module('itrak').factory('Equipment', function($resource, ServerName){
    return $resource(ServerName+'/equipment/:id', {id: '@id'});
  })

  // SubParts of a peice of Equipment
  angular.module('itrak').factory('SubParts', function($resource, ServerName){
    return $resource(ServerName+'/subparts/:id', {id: '@id'});
  })

  // Spares
  angular.module('itrak').factory('Spares', function($resource, ServerName){
    return $resource(ServerName+'/spares/:id', {id: '@id'});
  })

  // Consumables
  angular.module('itrak').factory('Consumables', function($resource, ServerName){
    return $resource(ServerName+'/consumables/:id', {id: '@id'});
  })

  // EquipTypes
  angular.module('itrak').factory('EquipTypes', function($resource, ServerName){
    return $resource(ServerName+'/equiptype/:id', {id: '@id'});
  })

  // WorkOrders
  angular.module('itrak').factory('WorkOrders', function($resource, ServerName){
    return $resource(ServerName+'/workorders/:id');
  })

  // Tasks
  angular.module('itrak').factory('Tasks', function($resource, ServerName){
    return $resource(ServerName+'/task/:id');
  })

  // Site Tasks
  angular.module('itrak').factory('SiteTasks', function($resource, ServerName){
    return $resource(ServerName+'/sitetask/:id', {id: '@id'});
  })

  // Site Equipment
  angular.module('itrak').factory('SiteEquipment', function($resource, ServerName){
    return $resource(ServerName+'/site_equipment/:id', {id: '@id'});
  })

  // Vendors
  angular.module('itrak').factory('Vendors', function($resource, ServerName){
    return $resource(ServerName+'/vendors/:id');
  })

  // Spare Parts
  angular.module('itrak').factory('Spares', function($resource, ServerName){
    return $resource(ServerName+'/spares/:id');
  })

  // Roles
  angular.module('itrak').factory('Roles', function($resource, ServerName){
    return $resource(ServerName+'/roles');
  })


})();
