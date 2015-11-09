(function(){
	// loginState
	'use strict';

	console.log("Loading LoginState")

	angular.module('itrak').service('loginState', function($state,$http){
		angular.extend(this, {
			loggedIn: false,
			username: '',
			toggle: function() {
				this.loggedIn = !this.loggedIn
				console.log("Login state toggled to", this.loggedIn)
			},
			login: function(username,passwd) { 
				console.log("logging in with credentials",
					username,
					passwd)

				var vm = this
				$http.post('http://localhost:8081/login', 
					{u: username, p: passwd})
					.then(function(){
						alert('logged in OK')
						vm.loggedIn = true
						vm.username = username
				        $state.go('home')
					},function(){
						alert('failed to login')
					})
			},
			logout: function() {
				this.loggedIn = false
			},
		})

		console.log('loginState init to',this)
	})

})();