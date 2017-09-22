"use strict"

import { DrawerNavigator } from "react-navigation"
import MapView from './subpages/mapview'
import ProfileView from './subpages/profile'

let mainView = DrawerNavigator({
    MapView: {
        path: "/",
        screen: MapView
    },
    ProfileView: {
        path: "/profile",
        screen: ProfileView
    }
})

export default mainView
