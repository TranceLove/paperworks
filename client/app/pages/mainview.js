"use strict"

import React from 'react';
import { TouchableOpacity } from 'react-native'
import { DrawerNavigator } from "react-navigation"
import { Icon } from 'react-native-elements'
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
},{
    drawerPosition: "left",
    drawerWidth: 200
})

mainView.navigationOptions = ({navigation}) => ({
    headerTitle: "Paperworks",
    headerLeft: (
        <TouchableOpacity style={{paddingHorizontal:20}} onPress={() => {
            if (navigation.state.index === 0) {
                // check if drawer is not open, then only open it
                navigation.navigate('DrawerOpen');
            } else {
                // else close the drawer
                navigation.navigate('DrawerClose');
            }
        }}>
            <Icon type="entypo" name="menu" size={24} />
        </TouchableOpacity>
    )
})

export default mainView
