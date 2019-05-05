import Vue from 'vue'
import Router from 'vue-router'
import bind from '@/bind/index'
import frame from '@/frame/index'
import projectFrame from '@/frame/projectFrame'
import newProject from '@/newProject/index'
import project from '@/project/index'
import mission from '@/mission/index'
import pace from '@/pace/index'
import projectDetail from '@/project/project'

Vue.use(Router)

export default new Router({
  routes: [
    {
      path: '/',
      name: 'bind',
      component: bind
    },
    {
      path: '/frame',
      name: 'frame',
      component: frame
    },
    {
      path: '/frame/:pName/:pId',
      name: 'projectFrame',
      component: projectFrame
    },
    {
      path: '/newProject',
      name: 'newProject',
      component: newProject
    },
    {
      path: '/project',
      name: 'project',
      component: project
    },
    {
      path: '/project/:pId',
      name: 'projectDetail',
      component: projectDetail
    },
    {
      path: '/mission',
      name: 'mission',
      component: mission
    },
    {
      path: '/pace',
      name: 'pace',
      component: pace
    }
  ]
})
