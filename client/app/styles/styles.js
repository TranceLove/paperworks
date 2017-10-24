"use strict"

import React from 'react';
import { StyleSheet } from 'react-native';

let selectServerStyle = StyleSheet.create({
    content: {
        flex:1,
        alignItems:'stretch',
        justifyContent:'center'
    },
    container: {
        height: 260,
        justifyContent:"space-between"
    },
    btnLocalhost: {
        marginBottom: 10,
        padding: 10
    },
    btnDevServer: {
        padding: 10
    },
    btnStagingServer: {
        padding: 10
    },
    btnProductionServer: {
        padding: 10
    }
})

let loginStyle = StyleSheet.create({
    content: {
        flex:1,
        alignItems:'stretch',
        justifyContent:'center'
    },
    container: {
        height: 180,
        justifyContent:"space-between"
    },
    twitterLogin: {
        marginBottom: 10,
        padding: 10
    },
    facebookLogin: {
        padding: 10,
        marginTop: 10
    }
})

let mainViewStyle = StyleSheet.create({
    map: {
        flex: 1
    }
})

export const Styles = {
    SelectServer: selectServerStyle,
    Login: loginStyle,
    MainView: mainViewStyle
}
