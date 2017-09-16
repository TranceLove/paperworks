"use strict"

import { Platform, StatusBar } from 'react-native';
import { StackNavigator } from "react-navigation"

import Login from "./pages/login"

export const Routes = StackNavigator({
    Login: {
        screen: Login,
        navigationOptions: {
            title: "Login",
            header: null
        }
    }
    // },
    // MainView: {
    //     screen: MainView,
    //     navigationOptions: {
    //         title: "Paperworks"
    //     }
    // },
    // Profile: {
    //     screen: Profile,
    //     navigationOptions: {
    //         title: "Profile"
    //     }
    // }
},{
    cardStyle: {
        paddingTop: Platform.OS === 'ios' ? 0 : StatusBar.currentHeight
    }
})
