"use strict"

import { Platform, StatusBar } from 'react-native';
import { StackNavigator } from "react-navigation"

import Login from "./pages/login"
import MainView from "./pages/mainview"

export const Routes = StackNavigator({
    Login: {
        screen: Login,
        navigationOptions: {
            title: "Login",
            header: null
        }
    },
    MainView: {
        screen: MainView,
        navigationOptions: {
            title: "Paperworks"
        }
    }
},{
    cardStyle: {
        paddingTop: Platform.OS === 'ios' ? 0 : StatusBar.currentHeight
    }
})
