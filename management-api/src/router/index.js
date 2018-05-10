import Vue from 'vue'
import Router from 'vue-router'
import Login from '@/components/Login'
import DashBoard from '@/components/DashBoard'
import Management from '@/components/Management'
import Admin from '@/components/Admin'
import User from '@/components/User'
import Api from '@/components/Api'
import Setting from '@/components/Setting'

Vue.use(Router)

export default new Router({
  routes: [
    {
      path: '/',
      name: 'Login',
      component: Login
    },
    {
      path: '/management',
      name: 'Management',
      component: Management,
      children: [
        { path: '', name: 'dashboard', component: DashBoard },
        { path: '/admin', name: 'admin', component: Admin },
        { path: '/user', name: 'user', component: User },
        { path: '/api', name: 'api', component: Api },
        { path: '/setting', name: 'setting', component: Setting }
      ]
    }
  ]
})
