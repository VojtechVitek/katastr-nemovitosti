// Code generated by wsdl2go. DO NOT EDIT.

package sestavysoapbinding

import (
	"github.com/fiorix/wsdl2go/soap"
)

// Namespace was auto-generated from WSDL.
var Namespace = "http://katastr.cuzk.cz/sestavy/v2.8"

// NewSestavy creates an initializes a Sestavy.
func NewSestavy(cli *soap.Client) Sestavy {
	return &sestavy{cli}
}

// Sestavy was auto-generated from WSDL
// and defines interface for the remote service. Useful for testing.
type Sestavy interface {
	// GenerujCenoveUdajeDleKu was auto-generated from WSDL.
	GenerujCenoveUdajeDleKu(GenerujCenoveUdajeDleKuRequest *GenerujCenoveUdajeDleKuRequest) (*GenerujSestavuResponse, error)

	// GenerujCenoveUdajeDleNemovitosti was auto-generated from WSDL.
	GenerujCenoveUdajeDleNemovitosti(GenerujCenoveUdajeDleNemovitostiRequest *GenerujCenoveUdajeDleNemovitostiRequest) (*GenerujSestavuResponse, error)

	// GenerujCenoveUdajeDleRizeni was auto-generated from WSDL.
	GenerujCenoveUdajeDleRizeni(GenerujCenoveUdajeDleRizeniRequest *GenerujCenoveUdajeDleRizeniRequest) (*GenerujSestavuResponse, error)

	// GenerujEvidenciPravProOsobu was auto-generated from WSDL.
	GenerujEvidenciPravProOsobu(GenerujEvidenciPravProOsobuRequest *GenerujEvidenciPravProOsobuRequest) (*GenerujSestavuResponse, error)

	// GenerujInfoOBodech was auto-generated from WSDL.
	GenerujInfoOBodech(GenerujInfoOBodechRequest *GenerujInfoOBodechRequest) (*GenerujSestavuResponse, error)

	// GenerujInfoOJednotkach was auto-generated from WSDL.
	GenerujInfoOJednotkach(GenerujInfoOJednotkachRequest *GenerujInfoOJednotkachRequest) (*GenerujSestavuResponse, error)

	// GenerujInfoOParcelach was auto-generated from WSDL.
	GenerujInfoOParcelach(GenerujInfoOParcelachRequest *GenerujInfoOParcelachRequest) (*GenerujSestavuResponse, error)

	// GenerujInfoOPravuStavby was auto-generated from WSDL.
	GenerujInfoOPravuStavby(GenerujInfoOPravuStavbyRequest *GenerujInfoOPravuStavbyRequest) (*GenerujSestavuResponse, error)

	// GenerujInfoORizeni was auto-generated from WSDL.
	GenerujInfoORizeni(GenerujInfoORizeniRequest *GenerujInfoORizeniRequest) (*GenerujSestavuResponse, error)

	// GenerujInfoOStavbach was auto-generated from WSDL.
	GenerujInfoOStavbach(GenerujInfoOStavbachRequest *GenerujInfoOStavbachRequest) (*GenerujSestavuResponse, error)

	// GenerujLV was auto-generated from WSDL.
	GenerujLV(GenerujLVRequest *GenerujLVRequest) (*GenerujSestavuResponse, error)

	// GenerujLVPresOS was auto-generated from WSDL.
	GenerujLVPresOS(GenerujLVPresOSRequest *GenerujLVPresOSRequest) (*GenerujSestavuResponse, error)

	// GenerujLVPresObjekty was auto-generated from WSDL.
	GenerujLVPresObjekty(GenerujLVPresObjektyRequest *GenerujLVPresObjektyRequest) (*GenerujSestavuResponse, error)

	// GenerujLVZjednodusene was auto-generated from WSDL.
	GenerujLVZjednodusene(GenerujLVZjednoduseneRequest *GenerujLVZjednoduseneRequest) (*GenerujSestavuResponse, error)

	// GenerujMapu was auto-generated from WSDL.
	GenerujMapu(GenerujMapuRequest *GenerujMapuRequest) (*GenerujSestavuResponse, error)

	// GenerujPrehledVlastnictvi was auto-generated from WSDL.
	GenerujPrehledVlastnictvi(GenerujPrehledVlastnictviRequest *GenerujPrehledVlastnictviRequest) (*GenerujSestavuResponse, error)

	// GenerujVystupZeSbirkyListin was auto-generated from WSDL.
	GenerujVystupZeSbirkyListin(GenerujVystupZeSbirkyListinRequest *GenerujVystupZeSbirkyListinRequest) (*GenerujSestavuResponse, error)

	// SeznamSestav was auto-generated from WSDL.
	SeznamSestav(SeznamSestavRequest *SeznamSestavRequest) (*GenerujSestavuResponse, error)

	// SmazSestavu was auto-generated from WSDL.
	SmazSestavu(SmazSestavuRequest *SmazSestavuRequest) (*GenerujSestavuResponse, error)

	// VratSestavu was auto-generated from WSDL.
	VratSestavu(VratSestavuRequest *VratSestavuRequest) (*GenerujSestavuResponse, error)

	// VypisUctu was auto-generated from WSDL.
	VypisUctu(VypisUctuRequest *VypisUctuRequest) (*GenerujSestavuResponse, error)
}

// Operation wrapper for GenerujCenoveUdajeDleKu.
// OperationGenerujCenoveUdajeDleKuRequest was auto-generated from
// WSDL.
type OperationGenerujCenoveUdajeDleKuRequest struct {
	Input *GenerujCenoveUdajeDleKuRequest `xml:"input,omitempty" json:"input,omitempty" yaml:"input,omitempty"`
}

// Operation wrapper for GenerujCenoveUdajeDleKu.
// OperationGenerujSestavuResponse was auto-generated from WSDL.
type OperationGenerujSestavuResponse struct {
	Output *GenerujSestavuResponse `xml:"output,omitempty" json:"output,omitempty" yaml:"output,omitempty"`
}

// Operation wrapper for GenerujCenoveUdajeDleNemovitosti.
// OperationGenerujCenoveUdajeDleNemovitostiRequest was auto-generated
// from WSDL.
type OperationGenerujCenoveUdajeDleNemovitostiRequest struct {
	Input *GenerujCenoveUdajeDleNemovitostiRequest `xml:"input,omitempty" json:"input,omitempty" yaml:"input,omitempty"`
}

// Operation wrapper for GenerujCenoveUdajeDleNemovitosti.
// OperationGenerujSestavuResponse was auto-generated from WSDL.
type OperationGenerujSestavuResponse struct {
	Output *GenerujSestavuResponse `xml:"output,omitempty" json:"output,omitempty" yaml:"output,omitempty"`
}

// Operation wrapper for GenerujCenoveUdajeDleRizeni.
// OperationGenerujCenoveUdajeDleRizeniRequest was auto-generated
// from WSDL.
type OperationGenerujCenoveUdajeDleRizeniRequest struct {
	Input *GenerujCenoveUdajeDleRizeniRequest `xml:"input,omitempty" json:"input,omitempty" yaml:"input,omitempty"`
}

// Operation wrapper for GenerujCenoveUdajeDleRizeni.
// OperationGenerujSestavuResponse was auto-generated from WSDL.
type OperationGenerujSestavuResponse struct {
	Output *GenerujSestavuResponse `xml:"output,omitempty" json:"output,omitempty" yaml:"output,omitempty"`
}

// Operation wrapper for GenerujEvidenciPravProOsobu.
// OperationGenerujEvidenciPravProOsobuRequest was auto-generated
// from WSDL.
type OperationGenerujEvidenciPravProOsobuRequest struct {
	Input *GenerujEvidenciPravProOsobuRequest `xml:"input,omitempty" json:"input,omitempty" yaml:"input,omitempty"`
}

// Operation wrapper for GenerujEvidenciPravProOsobu.
// OperationGenerujSestavuResponse was auto-generated from WSDL.
type OperationGenerujSestavuResponse struct {
	Output *GenerujSestavuResponse `xml:"output,omitempty" json:"output,omitempty" yaml:"output,omitempty"`
}

// Operation wrapper for GenerujInfoOBodech.
// OperationGenerujInfoOBodechRequest was auto-generated from WSDL.
type OperationGenerujInfoOBodechRequest struct {
	Input *GenerujInfoOBodechRequest `xml:"input,omitempty" json:"input,omitempty" yaml:"input,omitempty"`
}

// Operation wrapper for GenerujInfoOBodech.
// OperationGenerujSestavuResponse was auto-generated from WSDL.
type OperationGenerujSestavuResponse struct {
	Output *GenerujSestavuResponse `xml:"output,omitempty" json:"output,omitempty" yaml:"output,omitempty"`
}

// Operation wrapper for GenerujInfoOJednotkach.
// OperationGenerujInfoOJednotkachRequest was auto-generated from
// WSDL.
type OperationGenerujInfoOJednotkachRequest struct {
	Input *GenerujInfoOJednotkachRequest `xml:"input,omitempty" json:"input,omitempty" yaml:"input,omitempty"`
}

// Operation wrapper for GenerujInfoOJednotkach.
// OperationGenerujSestavuResponse was auto-generated from WSDL.
type OperationGenerujSestavuResponse struct {
	Output *GenerujSestavuResponse `xml:"output,omitempty" json:"output,omitempty" yaml:"output,omitempty"`
}

// Operation wrapper for GenerujInfoOParcelach.
// OperationGenerujInfoOParcelachRequest was auto-generated from
// WSDL.
type OperationGenerujInfoOParcelachRequest struct {
	Input *GenerujInfoOParcelachRequest `xml:"input,omitempty" json:"input,omitempty" yaml:"input,omitempty"`
}

// Operation wrapper for GenerujInfoOParcelach.
// OperationGenerujSestavuResponse was auto-generated from WSDL.
type OperationGenerujSestavuResponse struct {
	Output *GenerujSestavuResponse `xml:"output,omitempty" json:"output,omitempty" yaml:"output,omitempty"`
}

// Operation wrapper for GenerujInfoOPravuStavby.
// OperationGenerujInfoOPravuStavbyRequest was auto-generated from
// WSDL.
type OperationGenerujInfoOPravuStavbyRequest struct {
	Input *GenerujInfoOPravuStavbyRequest `xml:"input,omitempty" json:"input,omitempty" yaml:"input,omitempty"`
}

// Operation wrapper for GenerujInfoOPravuStavby.
// OperationGenerujInfoOPravuStavbyResponse was auto-generated
// from WSDL.
type OperationGenerujInfoOPravuStavbyResponse struct {
	Output *GenerujSestavuResponse `xml:"output,omitempty" json:"output,omitempty" yaml:"output,omitempty"`
}

// Operation wrapper for GenerujInfoORizeni.
// OperationGenerujInfoORizeniRequest was auto-generated from WSDL.
type OperationGenerujInfoORizeniRequest struct {
	Input *GenerujInfoORizeniRequest `xml:"input,omitempty" json:"input,omitempty" yaml:"input,omitempty"`
}

// Operation wrapper for GenerujInfoORizeni.
// OperationGenerujSestavuResponse was auto-generated from WSDL.
type OperationGenerujSestavuResponse struct {
	Output *GenerujSestavuResponse `xml:"output,omitempty" json:"output,omitempty" yaml:"output,omitempty"`
}

// Operation wrapper for GenerujInfoOStavbach.
// OperationGenerujInfoOStavbachRequest was auto-generated from
// WSDL.
type OperationGenerujInfoOStavbachRequest struct {
	Input *GenerujInfoOStavbachRequest `xml:"input,omitempty" json:"input,omitempty" yaml:"input,omitempty"`
}

// Operation wrapper for GenerujInfoOStavbach.
// OperationGenerujSestavuResponse was auto-generated from WSDL.
type OperationGenerujSestavuResponse struct {
	Output *GenerujSestavuResponse `xml:"output,omitempty" json:"output,omitempty" yaml:"output,omitempty"`
}

// Operation wrapper for GenerujLV.
// OperationGenerujLVRequest was auto-generated from WSDL.
type OperationGenerujLVRequest struct {
	Input *GenerujLVRequest `xml:"input,omitempty" json:"input,omitempty" yaml:"input,omitempty"`
}

// Operation wrapper for GenerujLV.
// OperationGenerujSestavuResponse was auto-generated from WSDL.
type OperationGenerujSestavuResponse struct {
	Output *GenerujSestavuResponse `xml:"output,omitempty" json:"output,omitempty" yaml:"output,omitempty"`
}

// Operation wrapper for GenerujLVPresOS.
// OperationGenerujLVPresOSRequest was auto-generated from WSDL.
type OperationGenerujLVPresOSRequest struct {
	Input *GenerujLVPresOSRequest `xml:"input,omitempty" json:"input,omitempty" yaml:"input,omitempty"`
}

// Operation wrapper for GenerujLVPresOS.
// OperationGenerujSestavuResponse was auto-generated from WSDL.
type OperationGenerujSestavuResponse struct {
	Output *GenerujSestavuResponse `xml:"output,omitempty" json:"output,omitempty" yaml:"output,omitempty"`
}

// Operation wrapper for GenerujLVPresObjekty.
// OperationGenerujLVPresObjektyRequest was auto-generated from
// WSDL.
type OperationGenerujLVPresObjektyRequest struct {
	Input *GenerujLVPresObjektyRequest `xml:"input,omitempty" json:"input,omitempty" yaml:"input,omitempty"`
}

// Operation wrapper for GenerujLVPresObjekty.
// OperationGenerujSestavuResponse was auto-generated from WSDL.
type OperationGenerujSestavuResponse struct {
	Output *GenerujSestavuResponse `xml:"output,omitempty" json:"output,omitempty" yaml:"output,omitempty"`
}

// Operation wrapper for GenerujLVZjednodusene.
// OperationGenerujLVZjednoduseneRequest was auto-generated from
// WSDL.
type OperationGenerujLVZjednoduseneRequest struct {
	Input *GenerujLVZjednoduseneRequest `xml:"input,omitempty" json:"input,omitempty" yaml:"input,omitempty"`
}

// Operation wrapper for GenerujLVZjednodusene.
// OperationGenerujSestavuResponse was auto-generated from WSDL.
type OperationGenerujSestavuResponse struct {
	Output *GenerujSestavuResponse `xml:"output,omitempty" json:"output,omitempty" yaml:"output,omitempty"`
}

// Operation wrapper for GenerujMapu.
// OperationGenerujMapuRequest was auto-generated from WSDL.
type OperationGenerujMapuRequest struct {
	Input *GenerujMapuRequest `xml:"input,omitempty" json:"input,omitempty" yaml:"input,omitempty"`
}

// Operation wrapper for GenerujMapu.
// OperationGenerujSestavuResponse was auto-generated from WSDL.
type OperationGenerujSestavuResponse struct {
	Output *GenerujSestavuResponse `xml:"output,omitempty" json:"output,omitempty" yaml:"output,omitempty"`
}

// Operation wrapper for GenerujPrehledVlastnictvi.
// OperationGenerujPrehledVlastnictviRequest was auto-generated
// from WSDL.
type OperationGenerujPrehledVlastnictviRequest struct {
	Input *GenerujPrehledVlastnictviRequest `xml:"input,omitempty" json:"input,omitempty" yaml:"input,omitempty"`
}

// Operation wrapper for GenerujPrehledVlastnictvi.
// OperationGenerujSestavuResponse was auto-generated from WSDL.
type OperationGenerujSestavuResponse struct {
	Output *GenerujSestavuResponse `xml:"output,omitempty" json:"output,omitempty" yaml:"output,omitempty"`
}

// Operation wrapper for GenerujVystupZeSbirkyListin.
// OperationGenerujVystupZeSbirkyListinRequest was auto-generated
// from WSDL.
type OperationGenerujVystupZeSbirkyListinRequest struct {
	Input *GenerujVystupZeSbirkyListinRequest `xml:"input,omitempty" json:"input,omitempty" yaml:"input,omitempty"`
}

// Operation wrapper for GenerujVystupZeSbirkyListin.
// OperationGenerujVystupZeSbirkyListinResponse was auto-generated
// from WSDL.
type OperationGenerujVystupZeSbirkyListinResponse struct {
	Output *GenerujSestavuResponse `xml:"output,omitempty" json:"output,omitempty" yaml:"output,omitempty"`
}

// Operation wrapper for SeznamSestav.
// OperationSeznamSestavRequest was auto-generated from WSDL.
type OperationSeznamSestavRequest struct {
	Input *SeznamSestavRequest `xml:"input,omitempty" json:"input,omitempty" yaml:"input,omitempty"`
}

// Operation wrapper for SeznamSestav.
// OperationGenerujSestavuResponse was auto-generated from WSDL.
type OperationGenerujSestavuResponse struct {
	Output *GenerujSestavuResponse `xml:"output,omitempty" json:"output,omitempty" yaml:"output,omitempty"`
}

// Operation wrapper for SmazSestavu.
// OperationSmazSestavuRequest was auto-generated from WSDL.
type OperationSmazSestavuRequest struct {
	Input *SmazSestavuRequest `xml:"input,omitempty" json:"input,omitempty" yaml:"input,omitempty"`
}

// Operation wrapper for SmazSestavu.
// OperationGenerujSestavuResponse was auto-generated from WSDL.
type OperationGenerujSestavuResponse struct {
	Output *GenerujSestavuResponse `xml:"output,omitempty" json:"output,omitempty" yaml:"output,omitempty"`
}

// Operation wrapper for VratSestavu.
// OperationVratSestavuRequest was auto-generated from WSDL.
type OperationVratSestavuRequest struct {
	Input *VratSestavuRequest `xml:"input,omitempty" json:"input,omitempty" yaml:"input,omitempty"`
}

// Operation wrapper for VratSestavu.
// OperationGenerujSestavuResponse was auto-generated from WSDL.
type OperationGenerujSestavuResponse struct {
	Output *GenerujSestavuResponse `xml:"output,omitempty" json:"output,omitempty" yaml:"output,omitempty"`
}

// Operation wrapper for VypisUctu.
// OperationVypisUctuRequest was auto-generated from WSDL.
type OperationVypisUctuRequest struct {
	Input *VypisUctuRequest `xml:"input,omitempty" json:"input,omitempty" yaml:"input,omitempty"`
}

// Operation wrapper for VypisUctu.
// OperationGenerujSestavuResponse was auto-generated from WSDL.
type OperationGenerujSestavuResponse struct {
	Output *GenerujSestavuResponse `xml:"output,omitempty" json:"output,omitempty" yaml:"output,omitempty"`
}

// sestavy implements the Sestavy interface.
type sestavy struct {
	cli *soap.Client
}

// GenerujCenoveUdajeDleKu was auto-generated from WSDL.
func (p *sestavy) GenerujCenoveUdajeDleKu(GenerujCenoveUdajeDleKuRequest *GenerujCenoveUdajeDleKuRequest) (*GenerujSestavuResponse, error) {
	α := struct {
		OperationGenerujCenoveUdajeDleKuRequest `xml:"tns:generujCenoveUdajeDleKu"`
	}{
		OperationGenerujCenoveUdajeDleKuRequest{
			GenerujCenoveUdajeDleKuRequest,
		},
	}

	γ := struct {
		OperationGenerujSestavuResponse `xml:"generujCenoveUdajeDleKuResponse"`
	}{}
	if err := p.cli.RoundTripWithAction("http://katastr.cuzk.cz/sestavy//generujCenoveUdajeDleKu", α, &γ); err != nil {
		return nil, err
	}
	return γ.GenerujSestavuResponse, nil
}

// GenerujCenoveUdajeDleNemovitosti was auto-generated from WSDL.
func (p *sestavy) GenerujCenoveUdajeDleNemovitosti(GenerujCenoveUdajeDleNemovitostiRequest *GenerujCenoveUdajeDleNemovitostiRequest) (*GenerujSestavuResponse, error) {
	α := struct {
		OperationGenerujCenoveUdajeDleNemovitostiRequest `xml:"tns:generujCenoveUdajeDleNemovitosti"`
	}{
		OperationGenerujCenoveUdajeDleNemovitostiRequest{
			GenerujCenoveUdajeDleNemovitostiRequest,
		},
	}

	γ := struct {
		OperationGenerujSestavuResponse `xml:"generujCenoveUdajeDleNemovitostiResponse"`
	}{}
	if err := p.cli.RoundTripWithAction("http://katastr.cuzk.cz/sestavy//generujCenoveUdajeDleNemovitosti", α, &γ); err != nil {
		return nil, err
	}
	return γ.GenerujSestavuResponse, nil
}

// GenerujCenoveUdajeDleRizeni was auto-generated from WSDL.
func (p *sestavy) GenerujCenoveUdajeDleRizeni(GenerujCenoveUdajeDleRizeniRequest *GenerujCenoveUdajeDleRizeniRequest) (*GenerujSestavuResponse, error) {
	α := struct {
		OperationGenerujCenoveUdajeDleRizeniRequest `xml:"tns:generujCenoveUdajeDleRizeni"`
	}{
		OperationGenerujCenoveUdajeDleRizeniRequest{
			GenerujCenoveUdajeDleRizeniRequest,
		},
	}

	γ := struct {
		OperationGenerujSestavuResponse `xml:"generujCenoveUdajeDleRizeniResponse"`
	}{}
	if err := p.cli.RoundTripWithAction("http://katastr.cuzk.cz/sestavy//generujCenoveUdajeDleRizeni", α, &γ); err != nil {
		return nil, err
	}
	return γ.GenerujSestavuResponse, nil
}

// GenerujEvidenciPravProOsobu was auto-generated from WSDL.
func (p *sestavy) GenerujEvidenciPravProOsobu(GenerujEvidenciPravProOsobuRequest *GenerujEvidenciPravProOsobuRequest) (*GenerujSestavuResponse, error) {
	α := struct {
		OperationGenerujEvidenciPravProOsobuRequest `xml:"tns:generujEvidenciPravProOsobu"`
	}{
		OperationGenerujEvidenciPravProOsobuRequest{
			GenerujEvidenciPravProOsobuRequest,
		},
	}

	γ := struct {
		OperationGenerujSestavuResponse `xml:"generujEvidenciPravProOsobuResponse"`
	}{}
	if err := p.cli.RoundTripWithAction("http://katastr.cuzk.cz/sestavy//generujEvidenciPravProOsobu", α, &γ); err != nil {
		return nil, err
	}
	return γ.GenerujSestavuResponse, nil
}

// GenerujInfoOBodech was auto-generated from WSDL.
func (p *sestavy) GenerujInfoOBodech(GenerujInfoOBodechRequest *GenerujInfoOBodechRequest) (*GenerujSestavuResponse, error) {
	α := struct {
		OperationGenerujInfoOBodechRequest `xml:"tns:generujInfoOBodech"`
	}{
		OperationGenerujInfoOBodechRequest{
			GenerujInfoOBodechRequest,
		},
	}

	γ := struct {
		OperationGenerujSestavuResponse `xml:"generujInfoOBodechResponse"`
	}{}
	if err := p.cli.RoundTripWithAction("http://katastr.cuzk.cz/sestavy//generujInfoOBodech", α, &γ); err != nil {
		return nil, err
	}
	return γ.GenerujSestavuResponse, nil
}

// GenerujInfoOJednotkach was auto-generated from WSDL.
func (p *sestavy) GenerujInfoOJednotkach(GenerujInfoOJednotkachRequest *GenerujInfoOJednotkachRequest) (*GenerujSestavuResponse, error) {
	α := struct {
		OperationGenerujInfoOJednotkachRequest `xml:"tns:generujInfoOJednotkach"`
	}{
		OperationGenerujInfoOJednotkachRequest{
			GenerujInfoOJednotkachRequest,
		},
	}

	γ := struct {
		OperationGenerujSestavuResponse `xml:"generujInfoOJednotkachResponse"`
	}{}
	if err := p.cli.RoundTripWithAction("http://katastr.cuzk.cz/sestavy//generujInfoOJednotkach", α, &γ); err != nil {
		return nil, err
	}
	return γ.GenerujSestavuResponse, nil
}

// GenerujInfoOParcelach was auto-generated from WSDL.
func (p *sestavy) GenerujInfoOParcelach(GenerujInfoOParcelachRequest *GenerujInfoOParcelachRequest) (*GenerujSestavuResponse, error) {
	α := struct {
		OperationGenerujInfoOParcelachRequest `xml:"tns:generujInfoOParcelach"`
	}{
		OperationGenerujInfoOParcelachRequest{
			GenerujInfoOParcelachRequest,
		},
	}

	γ := struct {
		OperationGenerujSestavuResponse `xml:"generujInfoOParcelachResponse"`
	}{}
	if err := p.cli.RoundTripWithAction("http://katastr.cuzk.cz/sestavy//generujInfoOParcelach", α, &γ); err != nil {
		return nil, err
	}
	return γ.GenerujSestavuResponse, nil
}

// GenerujInfoOPravuStavby was auto-generated from WSDL.
func (p *sestavy) GenerujInfoOPravuStavby(GenerujInfoOPravuStavbyRequest *GenerujInfoOPravuStavbyRequest) (*GenerujSestavuResponse, error) {
	α := struct {
		OperationGenerujInfoOPravuStavbyRequest `xml:"tns:generujInfoOPravuStavby"`
	}{
		OperationGenerujInfoOPravuStavbyRequest{
			GenerujInfoOPravuStavbyRequest,
		},
	}

	γ := struct {
		OperationGenerujInfoOPravuStavbyResponse `xml:"generujInfoOPravuStavbyResponse"`
	}{}
	if err := p.cli.RoundTripWithAction("http://katastr.cuzk.cz/sestavy//generujInfoOPravuStavby", α, &γ); err != nil {
		return nil, err
	}
	return γ.GenerujSestavuResponse, nil
}

// GenerujInfoORizeni was auto-generated from WSDL.
func (p *sestavy) GenerujInfoORizeni(GenerujInfoORizeniRequest *GenerujInfoORizeniRequest) (*GenerujSestavuResponse, error) {
	α := struct {
		OperationGenerujInfoORizeniRequest `xml:"tns:generujInfoORizeni"`
	}{
		OperationGenerujInfoORizeniRequest{
			GenerujInfoORizeniRequest,
		},
	}

	γ := struct {
		OperationGenerujSestavuResponse `xml:"generujInfoORizeniResponse"`
	}{}
	if err := p.cli.RoundTripWithAction("http://katastr.cuzk.cz/sestavy//generujInfoORizeni", α, &γ); err != nil {
		return nil, err
	}
	return γ.GenerujSestavuResponse, nil
}

// GenerujInfoOStavbach was auto-generated from WSDL.
func (p *sestavy) GenerujInfoOStavbach(GenerujInfoOStavbachRequest *GenerujInfoOStavbachRequest) (*GenerujSestavuResponse, error) {
	α := struct {
		OperationGenerujInfoOStavbachRequest `xml:"tns:generujInfoOStavbach"`
	}{
		OperationGenerujInfoOStavbachRequest{
			GenerujInfoOStavbachRequest,
		},
	}

	γ := struct {
		OperationGenerujSestavuResponse `xml:"generujInfoOStavbachResponse"`
	}{}
	if err := p.cli.RoundTripWithAction("http://katastr.cuzk.cz/sestavy//generujInfoOStavbach", α, &γ); err != nil {
		return nil, err
	}
	return γ.GenerujSestavuResponse, nil
}

// GenerujLV was auto-generated from WSDL.
func (p *sestavy) GenerujLV(GenerujLVRequest *GenerujLVRequest) (*GenerujSestavuResponse, error) {
	α := struct {
		OperationGenerujLVRequest `xml:"tns:generujLV"`
	}{
		OperationGenerujLVRequest{
			GenerujLVRequest,
		},
	}

	γ := struct {
		OperationGenerujSestavuResponse `xml:"generujLVResponse"`
	}{}
	if err := p.cli.RoundTripWithAction("http://katastr.cuzk.cz/sestavy//generujLV", α, &γ); err != nil {
		return nil, err
	}
	return γ.GenerujSestavuResponse, nil
}

// GenerujLVPresOS was auto-generated from WSDL.
func (p *sestavy) GenerujLVPresOS(GenerujLVPresOSRequest *GenerujLVPresOSRequest) (*GenerujSestavuResponse, error) {
	α := struct {
		OperationGenerujLVPresOSRequest `xml:"tns:generujLVPresOS"`
	}{
		OperationGenerujLVPresOSRequest{
			GenerujLVPresOSRequest,
		},
	}

	γ := struct {
		OperationGenerujSestavuResponse `xml:"generujLVPresOSResponse"`
	}{}
	if err := p.cli.RoundTripWithAction("http://katastr.cuzk.cz/sestavy//generujLVPresOS", α, &γ); err != nil {
		return nil, err
	}
	return γ.GenerujSestavuResponse, nil
}

// GenerujLVPresObjekty was auto-generated from WSDL.
func (p *sestavy) GenerujLVPresObjekty(GenerujLVPresObjektyRequest *GenerujLVPresObjektyRequest) (*GenerujSestavuResponse, error) {
	α := struct {
		OperationGenerujLVPresObjektyRequest `xml:"tns:generujLVPresObjekty"`
	}{
		OperationGenerujLVPresObjektyRequest{
			GenerujLVPresObjektyRequest,
		},
	}

	γ := struct {
		OperationGenerujSestavuResponse `xml:"generujLVPresObjektyResponse"`
	}{}
	if err := p.cli.RoundTripWithAction("http://katastr.cuzk.cz/sestavy//generujLVPresObjekty", α, &γ); err != nil {
		return nil, err
	}
	return γ.GenerujSestavuResponse, nil
}

// GenerujLVZjednodusene was auto-generated from WSDL.
func (p *sestavy) GenerujLVZjednodusene(GenerujLVZjednoduseneRequest *GenerujLVZjednoduseneRequest) (*GenerujSestavuResponse, error) {
	α := struct {
		OperationGenerujLVZjednoduseneRequest `xml:"tns:generujLVZjednodusene"`
	}{
		OperationGenerujLVZjednoduseneRequest{
			GenerujLVZjednoduseneRequest,
		},
	}

	γ := struct {
		OperationGenerujSestavuResponse `xml:"generujLVZjednoduseneResponse"`
	}{}
	if err := p.cli.RoundTripWithAction("http://katastr.cuzk.cz/sestavy//generujLVZjednodusene", α, &γ); err != nil {
		return nil, err
	}
	return γ.GenerujSestavuResponse, nil
}

// GenerujMapu was auto-generated from WSDL.
func (p *sestavy) GenerujMapu(GenerujMapuRequest *GenerujMapuRequest) (*GenerujSestavuResponse, error) {
	α := struct {
		OperationGenerujMapuRequest `xml:"tns:generujMapu"`
	}{
		OperationGenerujMapuRequest{
			GenerujMapuRequest,
		},
	}

	γ := struct {
		OperationGenerujSestavuResponse `xml:"generujMapuResponse"`
	}{}
	if err := p.cli.RoundTripWithAction("http://katastr.cuzk.cz/sestavy//generujMapu", α, &γ); err != nil {
		return nil, err
	}
	return γ.GenerujSestavuResponse, nil
}

// GenerujPrehledVlastnictvi was auto-generated from WSDL.
func (p *sestavy) GenerujPrehledVlastnictvi(GenerujPrehledVlastnictviRequest *GenerujPrehledVlastnictviRequest) (*GenerujSestavuResponse, error) {
	α := struct {
		OperationGenerujPrehledVlastnictviRequest `xml:"tns:generujPrehledVlastnictvi"`
	}{
		OperationGenerujPrehledVlastnictviRequest{
			GenerujPrehledVlastnictviRequest,
		},
	}

	γ := struct {
		OperationGenerujSestavuResponse `xml:"generujPrehledVlastnictviResponse"`
	}{}
	if err := p.cli.RoundTripWithAction("http://katastr.cuzk.cz/sestavy//generujPrehledVlastnictvi", α, &γ); err != nil {
		return nil, err
	}
	return γ.GenerujSestavuResponse, nil
}

// GenerujVystupZeSbirkyListin was auto-generated from WSDL.
func (p *sestavy) GenerujVystupZeSbirkyListin(GenerujVystupZeSbirkyListinRequest *GenerujVystupZeSbirkyListinRequest) (*GenerujSestavuResponse, error) {
	α := struct {
		OperationGenerujVystupZeSbirkyListinRequest `xml:"tns:generujVystupZeSbirkyListin"`
	}{
		OperationGenerujVystupZeSbirkyListinRequest{
			GenerujVystupZeSbirkyListinRequest,
		},
	}

	γ := struct {
		OperationGenerujVystupZeSbirkyListinResponse `xml:"generujVystupZeSbirkyListinResponse"`
	}{}
	if err := p.cli.RoundTripWithAction("http://katastr.cuzk.cz/sestavy//generujVystupZeSbirkyListin", α, &γ); err != nil {
		return nil, err
	}
	return γ.GenerujSestavuResponse, nil
}

// SeznamSestav was auto-generated from WSDL.
func (p *sestavy) SeznamSestav(SeznamSestavRequest *SeznamSestavRequest) (*GenerujSestavuResponse, error) {
	α := struct {
		OperationSeznamSestavRequest `xml:"tns:seznamSestav"`
	}{
		OperationSeznamSestavRequest{
			SeznamSestavRequest,
		},
	}

	γ := struct {
		OperationGenerujSestavuResponse `xml:"seznamSestavResponse"`
	}{}
	if err := p.cli.RoundTripWithAction("http://katastr.cuzk.cz/sestavy//seznamSestav", α, &γ); err != nil {
		return nil, err
	}
	return γ.GenerujSestavuResponse, nil
}

// SmazSestavu was auto-generated from WSDL.
func (p *sestavy) SmazSestavu(SmazSestavuRequest *SmazSestavuRequest) (*GenerujSestavuResponse, error) {
	α := struct {
		OperationSmazSestavuRequest `xml:"tns:smazSestavu"`
	}{
		OperationSmazSestavuRequest{
			SmazSestavuRequest,
		},
	}

	γ := struct {
		OperationGenerujSestavuResponse `xml:"smazSestavuResponse"`
	}{}
	if err := p.cli.RoundTripWithAction("http://katastr.cuzk.cz/sestavy//smazSestavu", α, &γ); err != nil {
		return nil, err
	}
	return γ.GenerujSestavuResponse, nil
}

// VratSestavu was auto-generated from WSDL.
func (p *sestavy) VratSestavu(VratSestavuRequest *VratSestavuRequest) (*GenerujSestavuResponse, error) {
	α := struct {
		OperationVratSestavuRequest `xml:"tns:vratSestavu"`
	}{
		OperationVratSestavuRequest{
			VratSestavuRequest,
		},
	}

	γ := struct {
		OperationGenerujSestavuResponse `xml:"vratSestavuResponse"`
	}{}
	if err := p.cli.RoundTripWithAction("http://katastr.cuzk.cz/sestavy//vratSestavu", α, &γ); err != nil {
		return nil, err
	}
	return γ.GenerujSestavuResponse, nil
}

// VypisUctu was auto-generated from WSDL.
func (p *sestavy) VypisUctu(VypisUctuRequest *VypisUctuRequest) (*GenerujSestavuResponse, error) {
	α := struct {
		OperationVypisUctuRequest `xml:"tns:vypisUctu"`
	}{
		OperationVypisUctuRequest{
			VypisUctuRequest,
		},
	}

	γ := struct {
		OperationGenerujSestavuResponse `xml:"vypisUctuResponse"`
	}{}
	if err := p.cli.RoundTripWithAction("http://katastr.cuzk.cz/sestavy//vypisUctu", α, &γ); err != nil {
		return nil, err
	}
	return γ.GenerujSestavuResponse, nil
}
