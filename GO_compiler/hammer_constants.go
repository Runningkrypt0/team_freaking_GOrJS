package main

const hammer_world_header string = `
versioninfo
{
	"editorversion" "400"
	"editorbuild" "6871"
	"mapversion" "1"
	"formatversion" "100"
	"prefab" "0"
}
visgroups
{
}
viewsettings
{
	"bSnapToGrid" "1"
	"bShowGrid" "1"
	"bShowLogicalGrid" "0"
	"nGridSpacing" "256"
	"bShow3DGrid" "0"
}
world
{
	"id" "1"
	"mapversion" "1"
	"classname" "worldspawn"
	"skyname" "sky_tf2_04"
	"maxpropscreenwidth" "-1"
	"detailvbsp" "detail_2fort.vbsp"
	"detailmaterial" "detail/detailsprites_2fort"`;
		
	
const hammer_world_footer string = `
entity
{
	"id" "2"
	"classname" "light_environment"
	"_ambient" "255 255 255 20"
	"_ambientHDR" "-1 -1 -1 1"
	"_AmbientScaleHDR" "1"
	"_light" "255 255 255 200"
	"_lightHDR" "-1 -1 -1 1"
	"_lightscaleHDR" "1"
	"angles" "0 0 0"
	"pitch" "-85"
	"SunSpreadAngle" "5"
	"origin" "545.735 464.133 9"
	editor
	{
		"color" "220 30 220"
		"visgroupshown" "1"
		"visgroupautoshown" "1"
		"logicalpos" "[0 500]"
	}
}
cameras
{
	"activecamera" "0"
	camera
	{
		"position" "[459.551 56.5543 390.069]"
		"look" "[446.428 53.6067 382.947]"
	}
}
cordons
{
	"mins" "(-128 -128 -1024)"
	"maxs" "(6592 6592 1024)"
	"active" "1"
}`;

const hammer_entity_header string = `
entity
{
	"id" "`;

const hammer_entity_footer string = `
	"classname" "prop_static"
	"angles" "0 75 0"
	"disableselfshadowing" "0"
	"disableshadows" "0"
	"disablevertexlighting" "0"
	"fademaxdist" "0"
	"fademindist" "-1"
	"fadescale" "1"
	"generatelightmaps" "0"
	"ignorenormals" "0"
	"lightmapresolutionx" "32"
	"lightmapresolutiony" "32"
	"maxdxlevel" "0"
	"mindxlevel" "0"
	"model" "models/props_badlands/barrel_flatbed01.mdl"
	"screenspacefade" "0"
	"skin" "0"
	"solid" "0"
	"origin" "1968 160 512"
	editor
	{
		"color" "255 255 0"
		"visgroupshown" "1"
		"visgroupautoshown" "1"
		"logicalpos" "[0 0]"
	}
}`;

const hammer_solid_header string = `
	solid
	{
		"id" "`;

const hammer_solid_footer string = `		editor
		{
			"color" "220 30 220"
			"visgroupshown" "1"
			"visgroupautoshown" "1"
		}
	}`;