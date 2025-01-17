/*
Candlepin

Candlepin is a subscription management server written in Java. It helps with management of software subscriptions.

API version: 4.4.15
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package caliri

import (
	"encoding/json"
)

// checks if the ProductDTO type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ProductDTO{}

// ProductDTO DTO representing the product data exposed to the API
type ProductDTO struct {
	Created *string `json:"created,omitempty"`
	Updated *string `json:"updated,omitempty"`
	Uuid *string `json:"uuid,omitempty"`
	Id *string `json:"id,omitempty"`
	Name *string `json:"name,omitempty"`
	Multiplier *int64 `json:"multiplier,omitempty"`
	Attributes []AttributeDTO `json:"attributes,omitempty"`
	ProductContent []ProductContentDTO `json:"productContent,omitempty"`
	DependentProductIds []string `json:"dependentProductIds,omitempty"`
	Branding []BrandingDTO `json:"branding,omitempty"`
	DerivedProduct *ProductDTO `json:"derivedProduct,omitempty"`
	ProvidedProducts []ProductDTO `json:"providedProducts,omitempty"`
	Href *string `json:"href,omitempty"`
}

// NewProductDTO instantiates a new ProductDTO object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewProductDTO() *ProductDTO {
	this := ProductDTO{}
	return &this
}

// NewProductDTOWithDefaults instantiates a new ProductDTO object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewProductDTOWithDefaults() *ProductDTO {
	this := ProductDTO{}
	return &this
}

// GetCreated returns the Created field value if set, zero value otherwise.
func (o *ProductDTO) GetCreated() string {
	if o == nil || IsNil(o.Created) {
		var ret string
		return ret
	}
	return *o.Created
}

// GetCreatedOk returns a tuple with the Created field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProductDTO) GetCreatedOk() (*string, bool) {
	if o == nil || IsNil(o.Created) {
		return nil, false
	}
	return o.Created, true
}

// HasCreated returns a boolean if a field has been set.
func (o *ProductDTO) HasCreated() bool {
	if o != nil && !IsNil(o.Created) {
		return true
	}

	return false
}

// SetCreated gets a reference to the given string and assigns it to the Created field.
func (o *ProductDTO) SetCreated(v string) {
	o.Created = &v
}

// GetUpdated returns the Updated field value if set, zero value otherwise.
func (o *ProductDTO) GetUpdated() string {
	if o == nil || IsNil(o.Updated) {
		var ret string
		return ret
	}
	return *o.Updated
}

// GetUpdatedOk returns a tuple with the Updated field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProductDTO) GetUpdatedOk() (*string, bool) {
	if o == nil || IsNil(o.Updated) {
		return nil, false
	}
	return o.Updated, true
}

// HasUpdated returns a boolean if a field has been set.
func (o *ProductDTO) HasUpdated() bool {
	if o != nil && !IsNil(o.Updated) {
		return true
	}

	return false
}

// SetUpdated gets a reference to the given string and assigns it to the Updated field.
func (o *ProductDTO) SetUpdated(v string) {
	o.Updated = &v
}

// GetUuid returns the Uuid field value if set, zero value otherwise.
func (o *ProductDTO) GetUuid() string {
	if o == nil || IsNil(o.Uuid) {
		var ret string
		return ret
	}
	return *o.Uuid
}

// GetUuidOk returns a tuple with the Uuid field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProductDTO) GetUuidOk() (*string, bool) {
	if o == nil || IsNil(o.Uuid) {
		return nil, false
	}
	return o.Uuid, true
}

// HasUuid returns a boolean if a field has been set.
func (o *ProductDTO) HasUuid() bool {
	if o != nil && !IsNil(o.Uuid) {
		return true
	}

	return false
}

// SetUuid gets a reference to the given string and assigns it to the Uuid field.
func (o *ProductDTO) SetUuid(v string) {
	o.Uuid = &v
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *ProductDTO) GetId() string {
	if o == nil || IsNil(o.Id) {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProductDTO) GetIdOk() (*string, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *ProductDTO) HasId() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *ProductDTO) SetId(v string) {
	o.Id = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *ProductDTO) GetName() string {
	if o == nil || IsNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProductDTO) GetNameOk() (*string, bool) {
	if o == nil || IsNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *ProductDTO) HasName() bool {
	if o != nil && !IsNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *ProductDTO) SetName(v string) {
	o.Name = &v
}

// GetMultiplier returns the Multiplier field value if set, zero value otherwise.
func (o *ProductDTO) GetMultiplier() int64 {
	if o == nil || IsNil(o.Multiplier) {
		var ret int64
		return ret
	}
	return *o.Multiplier
}

// GetMultiplierOk returns a tuple with the Multiplier field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProductDTO) GetMultiplierOk() (*int64, bool) {
	if o == nil || IsNil(o.Multiplier) {
		return nil, false
	}
	return o.Multiplier, true
}

// HasMultiplier returns a boolean if a field has been set.
func (o *ProductDTO) HasMultiplier() bool {
	if o != nil && !IsNil(o.Multiplier) {
		return true
	}

	return false
}

// SetMultiplier gets a reference to the given int64 and assigns it to the Multiplier field.
func (o *ProductDTO) SetMultiplier(v int64) {
	o.Multiplier = &v
}

// GetAttributes returns the Attributes field value if set, zero value otherwise.
func (o *ProductDTO) GetAttributes() []AttributeDTO {
	if o == nil || IsNil(o.Attributes) {
		var ret []AttributeDTO
		return ret
	}
	return o.Attributes
}

// GetAttributesOk returns a tuple with the Attributes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProductDTO) GetAttributesOk() ([]AttributeDTO, bool) {
	if o == nil || IsNil(o.Attributes) {
		return nil, false
	}
	return o.Attributes, true
}

// HasAttributes returns a boolean if a field has been set.
func (o *ProductDTO) HasAttributes() bool {
	if o != nil && !IsNil(o.Attributes) {
		return true
	}

	return false
}

// SetAttributes gets a reference to the given []AttributeDTO and assigns it to the Attributes field.
func (o *ProductDTO) SetAttributes(v []AttributeDTO) {
	o.Attributes = v
}

// GetProductContent returns the ProductContent field value if set, zero value otherwise.
func (o *ProductDTO) GetProductContent() []ProductContentDTO {
	if o == nil || IsNil(o.ProductContent) {
		var ret []ProductContentDTO
		return ret
	}
	return o.ProductContent
}

// GetProductContentOk returns a tuple with the ProductContent field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProductDTO) GetProductContentOk() ([]ProductContentDTO, bool) {
	if o == nil || IsNil(o.ProductContent) {
		return nil, false
	}
	return o.ProductContent, true
}

// HasProductContent returns a boolean if a field has been set.
func (o *ProductDTO) HasProductContent() bool {
	if o != nil && !IsNil(o.ProductContent) {
		return true
	}

	return false
}

// SetProductContent gets a reference to the given []ProductContentDTO and assigns it to the ProductContent field.
func (o *ProductDTO) SetProductContent(v []ProductContentDTO) {
	o.ProductContent = v
}

// GetDependentProductIds returns the DependentProductIds field value if set, zero value otherwise.
func (o *ProductDTO) GetDependentProductIds() []string {
	if o == nil || IsNil(o.DependentProductIds) {
		var ret []string
		return ret
	}
	return o.DependentProductIds
}

// GetDependentProductIdsOk returns a tuple with the DependentProductIds field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProductDTO) GetDependentProductIdsOk() ([]string, bool) {
	if o == nil || IsNil(o.DependentProductIds) {
		return nil, false
	}
	return o.DependentProductIds, true
}

// HasDependentProductIds returns a boolean if a field has been set.
func (o *ProductDTO) HasDependentProductIds() bool {
	if o != nil && !IsNil(o.DependentProductIds) {
		return true
	}

	return false
}

// SetDependentProductIds gets a reference to the given []string and assigns it to the DependentProductIds field.
func (o *ProductDTO) SetDependentProductIds(v []string) {
	o.DependentProductIds = v
}

// GetBranding returns the Branding field value if set, zero value otherwise.
func (o *ProductDTO) GetBranding() []BrandingDTO {
	if o == nil || IsNil(o.Branding) {
		var ret []BrandingDTO
		return ret
	}
	return o.Branding
}

// GetBrandingOk returns a tuple with the Branding field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProductDTO) GetBrandingOk() ([]BrandingDTO, bool) {
	if o == nil || IsNil(o.Branding) {
		return nil, false
	}
	return o.Branding, true
}

// HasBranding returns a boolean if a field has been set.
func (o *ProductDTO) HasBranding() bool {
	if o != nil && !IsNil(o.Branding) {
		return true
	}

	return false
}

// SetBranding gets a reference to the given []BrandingDTO and assigns it to the Branding field.
func (o *ProductDTO) SetBranding(v []BrandingDTO) {
	o.Branding = v
}

// GetDerivedProduct returns the DerivedProduct field value if set, zero value otherwise.
func (o *ProductDTO) GetDerivedProduct() ProductDTO {
	if o == nil || IsNil(o.DerivedProduct) {
		var ret ProductDTO
		return ret
	}
	return *o.DerivedProduct
}

// GetDerivedProductOk returns a tuple with the DerivedProduct field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProductDTO) GetDerivedProductOk() (*ProductDTO, bool) {
	if o == nil || IsNil(o.DerivedProduct) {
		return nil, false
	}
	return o.DerivedProduct, true
}

// HasDerivedProduct returns a boolean if a field has been set.
func (o *ProductDTO) HasDerivedProduct() bool {
	if o != nil && !IsNil(o.DerivedProduct) {
		return true
	}

	return false
}

// SetDerivedProduct gets a reference to the given ProductDTO and assigns it to the DerivedProduct field.
func (o *ProductDTO) SetDerivedProduct(v ProductDTO) {
	o.DerivedProduct = &v
}

// GetProvidedProducts returns the ProvidedProducts field value if set, zero value otherwise.
func (o *ProductDTO) GetProvidedProducts() []ProductDTO {
	if o == nil || IsNil(o.ProvidedProducts) {
		var ret []ProductDTO
		return ret
	}
	return o.ProvidedProducts
}

// GetProvidedProductsOk returns a tuple with the ProvidedProducts field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProductDTO) GetProvidedProductsOk() ([]ProductDTO, bool) {
	if o == nil || IsNil(o.ProvidedProducts) {
		return nil, false
	}
	return o.ProvidedProducts, true
}

// HasProvidedProducts returns a boolean if a field has been set.
func (o *ProductDTO) HasProvidedProducts() bool {
	if o != nil && !IsNil(o.ProvidedProducts) {
		return true
	}

	return false
}

// SetProvidedProducts gets a reference to the given []ProductDTO and assigns it to the ProvidedProducts field.
func (o *ProductDTO) SetProvidedProducts(v []ProductDTO) {
	o.ProvidedProducts = v
}

// GetHref returns the Href field value if set, zero value otherwise.
func (o *ProductDTO) GetHref() string {
	if o == nil || IsNil(o.Href) {
		var ret string
		return ret
	}
	return *o.Href
}

// GetHrefOk returns a tuple with the Href field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProductDTO) GetHrefOk() (*string, bool) {
	if o == nil || IsNil(o.Href) {
		return nil, false
	}
	return o.Href, true
}

// HasHref returns a boolean if a field has been set.
func (o *ProductDTO) HasHref() bool {
	if o != nil && !IsNil(o.Href) {
		return true
	}

	return false
}

// SetHref gets a reference to the given string and assigns it to the Href field.
func (o *ProductDTO) SetHref(v string) {
	o.Href = &v
}

func (o ProductDTO) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ProductDTO) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Created) {
		toSerialize["created"] = o.Created
	}
	if !IsNil(o.Updated) {
		toSerialize["updated"] = o.Updated
	}
	if !IsNil(o.Uuid) {
		toSerialize["uuid"] = o.Uuid
	}
	if !IsNil(o.Id) {
		toSerialize["id"] = o.Id
	}
	if !IsNil(o.Name) {
		toSerialize["name"] = o.Name
	}
	if !IsNil(o.Multiplier) {
		toSerialize["multiplier"] = o.Multiplier
	}
	if !IsNil(o.Attributes) {
		toSerialize["attributes"] = o.Attributes
	}
	if !IsNil(o.ProductContent) {
		toSerialize["productContent"] = o.ProductContent
	}
	if !IsNil(o.DependentProductIds) {
		toSerialize["dependentProductIds"] = o.DependentProductIds
	}
	if !IsNil(o.Branding) {
		toSerialize["branding"] = o.Branding
	}
	if !IsNil(o.DerivedProduct) {
		toSerialize["derivedProduct"] = o.DerivedProduct
	}
	if !IsNil(o.ProvidedProducts) {
		toSerialize["providedProducts"] = o.ProvidedProducts
	}
	if !IsNil(o.Href) {
		toSerialize["href"] = o.Href
	}
	return toSerialize, nil
}

type NullableProductDTO struct {
	value *ProductDTO
	isSet bool
}

func (v NullableProductDTO) Get() *ProductDTO {
	return v.value
}

func (v *NullableProductDTO) Set(val *ProductDTO) {
	v.value = val
	v.isSet = true
}

func (v NullableProductDTO) IsSet() bool {
	return v.isSet
}

func (v *NullableProductDTO) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableProductDTO(val *ProductDTO) *NullableProductDTO {
	return &NullableProductDTO{value: val, isSet: true}
}

func (v NullableProductDTO) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableProductDTO) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


