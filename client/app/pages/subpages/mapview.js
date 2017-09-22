"use strict"

import React from 'react';
import MapView from 'react-native-maps';
import { Styles } from './../../styles/styles'
import ProfileService from './../../services/ProfileService'

export default ({navigation}) => (
    <MapView style={Styles.MainView.map}
        region={ProfileService.myLastLocation()}
        showsCompass={true}
        showsScale={true}
        loadingEnabled={true}
        onRegionChange={(region)=> console.log(region.latitudeDelta + ", " + region.longitudeDelta)}
    />
)
