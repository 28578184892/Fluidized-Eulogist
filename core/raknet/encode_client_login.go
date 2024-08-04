package RaknetConnection

import (
	fbauth "Eulogist/core/fb_auth/pv4"
	"Eulogist/core/minecraft/protocol"
	"Eulogist/core/minecraft/protocol/login"
	"bytes"
	"crypto/ecdsa"
	"encoding/base64"
	"fmt"
	"math/rand"
	"time"

	"github.com/google/uuid"
)

// EncodeLogin 编码登录请求。
// 它使用提供的身份验证响应、
// 客户端密钥和皮肤信息生成登录请求数据包
func (r *Raknet) EncodeLogin(
	authResponse fbauth.AuthResponse,
	clientKey *ecdsa.PrivateKey,
	skin *Skin,
) ([]byte, error) {
	identityData := login.IdentityData{}
	clientData := login.ClientData{}

	// 设置默认的身份数据
	defaultIdentityData(&identityData)
	// 设置默认的客户端数据
	err := defaultClientData(&clientData, authResponse, skin)
	if err != nil {
		return nil, fmt.Errorf("EncodeLogin: %v", err)
	}

	var request []byte
	// 我们以 Android 设备登录，这将在 JWT 链中的 titleId 字段中显示。
	// 这些字段无法被编辑，而我们也仅仅是强制以 Android 数据进行登录
	setAndroidData(&clientData)

	// 编码登录请求
	request = login.Encode(authResponse.ChainInfo, clientData, clientKey)
	// 解析身份数据以确保其有效
	identityData, _, _, err = login.Parse(request)
	if err != nil {
		return nil, fmt.Errorf("EncodeLogin: WARNING: Identity data parsing error: %v", err)
	}

	return request, nil
}

// defaultIdentityData 编辑传入的 IdentityData，
// 为所有未更改的字段设置默认值
func defaultIdentityData(data *login.IdentityData) {
	if data.Identity == "" {
		data.Identity = uuid.New().String()
	}
	if data.DisplayName == "" {
		data.DisplayName = "Steve"
	}
}

// defaultClientData 编辑传入的 ClientData，
// 为所有未更改的字段设置默认值
func defaultClientData(
	d *login.ClientData,
	authResponse fbauth.AuthResponse,
	skin *Skin,
) error {
	rand.Seed(time.Now().Unix())

	d.ServerAddress = authResponse.RentalServerIP
	d.ThirdPartyName = authResponse.BotName
	if d.DeviceOS == 0 {
		d.DeviceOS = protocol.DeviceAndroid
	}
	if d.GameVersion == "" {
		d.GameVersion = protocol.CurrentVersion
	}

	// PhoenixBuilder specific changes.
	// Author: Liliya233, Happy2018new
	if d.GrowthLevel != authResponse.BotLevel {
		d.GrowthLevel = authResponse.BotLevel
	}

	if d.ClientRandomID == 0 {
		d.ClientRandomID = rand.Int63()
	}
	if d.DeviceID == "" {
		d.DeviceID = uuid.New().String()
	}
	if d.LanguageCode == "" {
		// PhoenixBuilder specific changes.
		// Author: Liliya233
		d.LanguageCode = "zh_CN"
		// d.LanguageCode = "en_GB"
	}
	if d.AnimatedImageData == nil {
		d.AnimatedImageData = make([]login.SkinAnimation, 0)
	}
	if d.PersonaPieces == nil {
		d.PersonaPieces = make([]login.PersonaPiece, 0)
	}
	if d.PieceTintColours == nil {
		d.PieceTintColours = make([]login.PersonaPieceTintColour, 0)
	}
	if d.SelfSignedID == "" {
		d.SelfSignedID = uuid.New().String()
	}
	if d.SkinID == "" {
		d.SkinID = uuid.New().String()
	}
	if d.SkinItemID == "" {
		d.SkinItemID = authResponse.SkinInfo.ItemID
	}
	if d.SkinData == "" {
		if skin != nil {
			d.SkinData = base64.StdEncoding.EncodeToString(skin.SkinPixels)
			d.SkinImageHeight, d.SkinImageWidth = skin.SkinHight, skin.SkinWidth
			d.SkinGeometry = base64.StdEncoding.EncodeToString(skin.SkinGeometry)
			d.SkinGeometryVersion = base64.StdEncoding.EncodeToString([]byte("0.0.0"))
			d.SkinResourcePatch = base64.StdEncoding.EncodeToString(skin.SkinResourcePatch)
			d.PremiumSkin = true
		} else {
			d.SkinData = base64.StdEncoding.EncodeToString(bytes.Repeat([]byte{0, 0, 0, 255}, 32*64))
			d.SkinImageHeight = 32
			d.SkinImageWidth = 64
		}
	}
	if d.SkinResourcePatch == "" {
		d.SkinResourcePatch = base64.StdEncoding.EncodeToString(defaultSkinResourcePatch)
	}
	if d.SkinGeometry == "" {
		d.SkinGeometry = base64.StdEncoding.EncodeToString(defaultSkinGeometry)
	}

	return nil
}

// setAndroidData 确保传入的 login.ClientData
// 匹配您在 Android 设备上看到的设置
func setAndroidData(data *login.ClientData) {
	data.DeviceOS = protocol.DeviceAndroid
	data.GameVersion = protocol.CurrentVersion
}
