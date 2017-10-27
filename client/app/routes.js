"use strict"

import { Platform, StatusBar } from 'react-native';
import { StackNavigator } from "react-navigation"

import Login from "./pages/login"
import MainView from "./pages/mainview"
import SelectServer from "./pages/selectserver"

const isDev = true

export const Routes = StackNavigator({
    SelectServer: {
        screen: SelectServer,
        navigationOptions: {
            title: "Select Server",
            header: null
        }
    },
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
    initialRouteName: isDev ? "SelectServer" : "Login",
    cardStyle: {
        paddingTop: Platform.OS === 'ios' ? 0 : StatusBar.currentHeight
    }
})
