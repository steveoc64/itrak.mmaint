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
    return $resource(ServerName+'/site/:id');
  })

  // Equipment
  angular.module('itrak').factory('Equipment', function($resource, ServerName){
    return $resource(ServerName+'/equipment/:id');
  })

  // WorkOrders
  angular.module('itrak').factory('WorkOrders', function($resource, ServerName){
    return $resource(ServerName+'/workorders/:id');
  })

  // Tasks
  angular.module('itrak').factory('Tasks', function($resource, ServerName){
    return $resource(ServerName+'/task/:id');
  })

  // Vendors
  angular.module('itrak').factory('Vendors', function($resource, ServerName){
    return $resource(ServerName+'/vendor/:id');
  })

  // Spare Parts
  angular.module('itrak').factory('Spares', function($resource, ServerName){
    return $resource(ServerName+'/spares/:id');
  })

  // Roles
  angular.module('itrak').factory('Roles', function($resource, ServerName){
    return $resource(ServerName+'/roles');
  })

  // Vendors
  angular.module('itrak').factory('CiboLate', function($resource, ServerName){
    return $resource(ServerName+'/Cibo');
  })



})();
