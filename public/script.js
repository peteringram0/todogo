new Vue({
  el: '#vue',
  data: {
    token: false,
    user: false,
    link: false,
    task: {
      name: null
    }
  },
  methods: {
    login() {
      window.location.href = this.link;
    },
    logout() {
      window.axios.get('/api/v1/logout',{headers: {"Authorization" : "Bearer " + this.token}})
        .then(function(r) {
          if(r.status === 200) {
            localStorage.removeItem('token');
            location.reload();
          }
        }.bind(this))
    },
    getGoogleLink() {
      window.axios.get('/api/v1/login')
        .then(function(r) {
          this.link = r.data.link
        }.bind(this))
    },
    getMe() {
      window.axios.get('/api/v1/me',{headers: {"Authorization" : "Bearer " + this.token}})
        .then(function(r) {

          this.user = r.data;

          // Display it
          this.$refs.vue.style = "display: block;"

        }.bind(this))
    },
    addTask() {
      window.axios.post('/api/v1/tasks', {
        name: this.task.name
      }, {headers: {"Authorization" : "Bearer " + this.token}})
        .then(function(r) {
          if(r.status === 201) {

            if(this.user.tasks.items === null)
              this.user.tasks.items = [];

            this.user.tasks.items.push(r.data);

            this.task.name = null;

          }
        }.bind(this))
    },
    edit(item) {
      window.axios.put('/api/v1/tasks/' + item.id, {
        name: item.name
      }, {headers: {"Authorization" : "Bearer " + this.token}})
    },
    deleteItem(id, index) {
      window.axios.delete('/api/v1/tasks/' + id, {headers: {"Authorization" : "Bearer " + this.token}})
        .then(function(r) {
          console.log(r)
          if(r.status == 200)
            this.user.tasks.items.splice(index, 1);
        }.bind(this))
    }
  },
  mounted() {

    // Get token from URL param
    var tokenInURL = getParameterByName('token');
    var tokenInStorage = localStorage.getItem('token');

    // If there ia a token in the URL and no token in storage
    if(tokenInURL && tokenInStorage === null) {
      localStorage.setItem("token", tokenInURL);
      window.history.replaceState(null, null, window.location.pathname);
      window.location.href = window.location.pathname;
    } else if(tokenInStorage !== null) {
      this.token = tokenInStorage;
      this.getMe();
    } else if(!tokenInURL && tokenInStorage === null) {
      this.getGoogleLink();
    }

  }
})

function getParameterByName(name, url) {
  if (!url) url = window.location.href;
  name = name.replace(/[\[\]]/g, "\\$&");
  var regex = new RegExp("[?&]" + name + "(=([^&#]*)|&|#|$)"),
      results = regex.exec(url);
  if (!results) return false;
  if (!results[2]) return '';
  return decodeURIComponent(results[2].replace(/\+/g, " "));
}
