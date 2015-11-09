(function(){
	// loginState
	'use strict';

	console.log("Loading LoginState")

	angular.module('itrak').service('loginState', function($state,$http){
		angular.extend(this, {
			loggedIn: false,
			toggle: function() {
				this.loggedIn = !this.loggedIn
				console.log("Login state toggled to", this.loggedIn)
			},
			login: function(username,passwd) { 
				console.log("logging in with credentials",
					username,
					passwd)

				$http.put('http://localhost:8081/login', {u: username, p: passwd})
				this.loggedIn = true
		        $state.go('home')
			},
			logout: function() {
				this.loggedIn = false
			},
		})

		console.log('loginState init to',this)
	})

})();